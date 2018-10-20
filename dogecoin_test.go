package dogecoin

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/alivanz/go-crypto/bitcoin"
)

func TestGetUnspentData(t *testing.T) {
	// mendefinisikan private key awal yang akan digunakan

	/*address : DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR*/
	// privkey := WIFtoPrivKey("QQ5bFuoytx3CawBrsZeXNkW32MPxQvd2W95QeCqk13G7yWYecJoD")

	// /*address : DUJ4EoqvPAzCpJBACQa6vYuQpcv9FSHF5d*/
	// privkey := WIFtoPrivKey("QVBG5ywrW8nxfZWEFK7q9Jgi55SqjiMf1Fs275VLhd76Ta9VhwYp")

	// /*address : DGaL1Bm3YmWCnuz4j3BSPQZSNYGysgX9ZL*/
	// privkey := "e105500a65cd0eda7ec6784a27a09f20c725ade74ec7d1bd96d09318d0ed43a4"

	// /*address : DK9m4nYAaZHoYRkW7eZzjNLTgHdjpMEmkm*/
	privkey := "3cd0560f5b27591916c643a0b7aa69d03839380a738d2e912990dcc573715d2c"

	// konversi private key ke object wallet yang berisi pasangan private key dan public key
	privkeybin, err := hex.DecodeString(privkey)
	ErrorCheck(err)
	wallet, err := bitcoin.NewWallet(privkeybin)
	ErrorCheck(err)

	// membentuk pubkeyhash dari object wallet
	pubkeyhash := WalletToPubKeyHash(wallet)

	// membentuk data coin cryptocurrency dari pubkeyhash
	var coin Doge
	coindata := coin.CreateCoin(pubkeyhash)

	// menampilkan data address dan saldo user
	log.Printf("myaddress    : %v\n", coindata.address)
	log.Printf("balance      : %v\n", coindata.balance)

	// menampilkan list data output dari transaksi lain yang dimiliki oleh user yang masih belum dipakai dari value terbesar hingga terkecil
	for _, row := range coindata.unspent {
		log.Printf("%v %v %v", row.TxHash, row.TxOutputN, row.Value)
	}

	// membuat array dari tipe bentukan "Destination", yang berisi address tujuan dan jumlah yang ingin dikirimkan
	dest := make([]Destination, 0)

	// // menambah elemen array dari address tujuan yang telah didefinisikan sebelumnya
	// dest = append(dest, Destination{"DUJ4EoqvPAzCpJBACQa6vYuQpcv9FSHF5d", uint64(450000000)})
	// dest = append(dest, Destination{"DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR", uint64(300000000)})

	//splitter
	dest = append(dest, Destination{coindata.address, uint64(900000000)})
	dest = append(dest, Destination{coindata.address, uint64(650000000)})

	// menampilkan list address tujuan dan jumlah yang dikirimkan
	var sendvalue uint64
	for _, outaddr := range dest {
		log.Printf("Destination : %v , Value : %v\n", outaddr.Address, outaddr.Value)
		sendvalue = sendvalue + outaddr.Value
	}

	// menghitung total fee dan jumlah list output yang dipakai untuk pengiriman sejumlah doge yang telah didefinisikan
	// serta menambah elemen array destinasi yang baru ke address user jika ada kembalian doge yang perlu dibayar ke user
	totalfee, numindex := ChangeUnspent(coindata, sendvalue, &dest)

	// menampilkan total fee dan jumlah unspent yg dipakaidari proses transaksi
	log.Printf("totalfee     : %v\n", IntToStr(totalfee))
	log.Printf("used_unspent : %v\n", numindex)

	// mengecek apakah jumlah saldo masih mencukupi, jika iya, proses pembuatan hex transaksi dijalankan
	if coindata.balance >= (sendvalue + totalfee) {

		// membuat hex transaksi yang sudah di signature
		signtx := CreateSignedTransaction(coindata, dest, wallet, numindex)

		// menampilkan hex transaksi yang sudah di signature di console
		log.Printf("signtxhex    : %v\n", signtx)

		// broadcast transaksi
		ErrorCheck(coin.Broadcast(signtx))

		// jika jumlah saldo tidak mencukupi jumlah pembayaran beserta total fee
	} else if coindata.balance >= sendvalue {
		log.Printf("total fee belum melewati batas minimum, transaksi tidak dapat dilakukan\n")

		// jika jumlah saldo lebih kecil dari jumlah pembayaran
	} else {
		log.Printf("saldo tidak mencukupi, transaksi tidak dapat dilakukan\n")
	}
}
