package dogecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	crypto "github.com/alivanz/go-crypto"
	"github.com/alivanz/go-crypto/bitcoin/base58"
)

const (
	sighashcode = "01"
	header      = "30"
	integer     = "02"
)

type Destination struct {
	Address string
	Value   uint64
}

// mengurutkan list data output dari transaksi lain yang dimiliki address dari value terbesar hingga yang terkecil
func OrderUnspent(unspent *[]Unspent) {
	var temp Unspent
	for i := 1; i < len(*unspent); i++ {
		for j := i; (j > 0) && ((*unspent)[j].Value > (*unspent)[j-1].Value); j-- {
			temp = (*unspent)[j]
			(*unspent)[j] = (*unspent)[j-1]
			(*unspent)[j-1] = temp
		}
	}
}

// menghitung jumlah list output yang dipakai untuk pengiriman sejumlah doge yang telah didefinisikan di "sendvalue"
// serta menambah elemen array destinasi "dest" yang baru ke address pengirim jika ada kembalian doge yang perlu dibayar ke address pengirim
func ChangeUnspent(coindata Coin, sendvalue uint64, dest *[]Destination) (uint64, int) {
	var i int
	outfee := coindata.Fee
	infee := 2 * outfee
	totalfee := outfee * uint64(len(*dest))
	sending := sendvalue + totalfee
	for i = 0; i < len(coindata.Unspent); i++ {
		totalfee = totalfee + infee
		if coindata.Unspent[i].Value > infee {
			switch {
			case sending >= (coindata.Unspent[i].Value - infee):
				sending = sending - (coindata.Unspent[i].Value - infee)
				if sending == 0 {
					i++
					return totalfee, i
				}
			case sending >= (coindata.Unspent[i].Value - (infee + outfee)):
				totalfee = totalfee + (coindata.Unspent[i].Value - sending) - infee
				i++
				return totalfee, i
			default:
				totalfee = totalfee + outfee
				change := (coindata.Unspent[i].Value - (infee + outfee)) - sending
				(*dest) = append((*dest), Destination{coindata.Address, change})
				i++
				return totalfee, i
			}
		}
	}
	return totalfee, i
}

// membuat format hex untuk bagian input, baik untuk hex transaksi yang masih raw (belum di-signature) ataupun yang sudah di-signature
/* pengecekan apakah input object walet terdeteksi pada fungsi atau tidak menentukan apakah fungsi input template membuat input hex
untuk hex transaksi raw atau hex transaksi yang di-signature*/
// numindex menunjukkan jumlah list output yang dipakai untuk proses transaksi, untuk selanjutnya bisa disebut sebagai jumlah input
/* posindex menunjukkan pada input ke-berapa scriptpubkey akan diletakkan pada posisi scriptsig pada hex transaksi raw
yang mana selanjutnya hanya input pada posindex tersebut yang akan di-signature*/
func InputTemplate(coindata Coin, dest []Destination, wallet crypto.Wallet, numindex int, posindex int) string {
	var input, inputfinal bytes.Buffer
	var i int
	var scriptsig, index string
	for i = 0; i < numindex; i++ {
		// hash transaksi sebelumnya
		input.WriteString(ReverseHex(coindata.Unspent[i].TxHash))
		// index input (atau output yang tidak dipakai pada transaksi sebelumnya)
		index = fmt.Sprintf("%x", coindata.Unspent[i].TxOutputN)
		for len(index) < 8 {
			index = "0" + index
		}
		// index input dibuat ke dalam format little endian
		input.WriteString(ReverseHex(index))
		// menentukan pada posisi hex scriptsig, yang diletakkan apakah signature atau scriptpubkey
		switch {
		// jika input wallet terdeteksi, signature diletakkan pada posisi scriptsig di hex transaksi
		case wallet != nil:
			scriptsig = CreateSignature(coindata, dest, wallet, numindex, i)
			input.WriteString(VarInt(len(scriptsig) / 2))
			input.WriteString(scriptsig)
		// jika input wallet tidak terdeteksi, namun index input yang dipakai sama dengan posindex,
		// maka scriptpubkey diletakkan pada input tersebut
		case (wallet == nil) && (i == posindex):
			input.WriteString(VarInt(len(coindata.Unspent[i].Script) / 2))
			input.WriteString(coindata.Unspent[i].Script)
		// kondisi jika input wallet tidak terdeteksi dan juga index input masih belum sesuai dengan posindex, scriptsig dibiarkan kosong
		default:
			input.WriteString("00")
		}
		input.WriteString("ffffffff")
	}
	// jumlah input yang dipakai
	inputfinal.WriteString(VarInt(i))
	// pembentukan final hex transaksi bagian input
	inputfinal.WriteString(input.String())
	return inputfinal.String()
}

func OutputTemplate(dest []Destination) string {
	var output, outputfinal bytes.Buffer
	var i int
	var value, pubkeyhash, scriptpubkey string
	for i = 0; i < len(dest); i++ {
		// menentukan jumlah value tertentu yang ingin dikirimkan ke address tertentu
		value = fmt.Sprintf("%x", dest[i].Value)
		for len(value) < 16 {
			value = "0" + value
		}
		// index output dibuat ke dalam format little endian
		output.WriteString(ReverseHex(value))
		// decode address tujuan ke dalam tipe byte
		binaddress, err := base58.Decode(dest[i].Address)
		ErrorCheck(err)
		// mengekstrak pubkeyhash dari address tujuan yang sudah di-decode
		pubkeyhash = BinAddressPubKeyHash(hex.EncodeToString(binaddress))
		// pembentukan scriptpukey untuk address tujuan untuk mode transaksi pay-to-pubkey-hash
		scriptpubkey = P2PKH(pubkeyhash)
		output.WriteString(VarInt(len(scriptpubkey) / 2))
		output.WriteString(scriptpubkey)
	}
	// jumlah output atau address tujuan
	outputfinal.WriteString(VarInt(i))
	// pembentukan final hex transaksi bagian output
	outputfinal.WriteString(output.String())
	return outputfinal.String()
}

// pembentukan signature untuk diletakkan pada posisi scriptsig pada hex transaksi
func ScriptSig(r, s, pubkey string) string {
	var sign, signfinal, scriptsig bytes.Buffer

	sign.WriteString(integer)
	sign.WriteString(VarInt(len(r) / 2))
	sign.WriteString(r)
	sign.WriteString(integer)
	sign.WriteString(VarInt(len(s) / 2))
	sign.WriteString(s)

	signfinal.WriteString(header)
	signfinal.WriteString(VarInt(len(sign.String()) / 2))
	signfinal.WriteString(sign.String())
	signfinal.WriteString(sighashcode)

	scriptsig.WriteString(VarInt(len(signfinal.String()) / 2))
	scriptsig.WriteString(signfinal.String())
	scriptsig.WriteString(VarInt(len(pubkey) / 2))
	scriptsig.WriteString(pubkey)

	return scriptsig.String()
}
