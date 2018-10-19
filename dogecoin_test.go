package dogecoin

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/alivanz/go-crypto/bitcoin"
)

func TestGetUnspentData(t *testing.T) {
	// mendefinisikan private key awal yang akan digunakan
	privkey := "e105500a65cd0eda7ec6784a27a09f20c725ade74ec7d1bd96d09318d0ed43a4"
	// konversi private key ke object wallet yang berisi pasangan private key dan public key
	privkeybin, _ := hex.DecodeString(privkey)
	wallet, _ := bitcoin.NewWallet(privkeybin)
	// membentuk address dari public key yang diperoleh dari object wallet
	myaddress := WalletToAddress(wallet)
	// address tujuan dan jumlah yang ingin dikirimkan
	destaddress := "DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR"
	sendvalue := uint64(2362500000)
	// mengambil data saldo dari address user
	balance := GetBalance(myaddress)
	// menampilkan data address dan saldo user, serta address tujuan dan destinasi di console
	log.Printf("myaddress    : %v\n", myaddress)
	log.Printf("balance      : %v\n", balance.Balance)
	log.Printf("destination  : %v\n", destaddress)
	log.Printf("sendvalue    : %v\n", IntToStr(sendvalue))
	// mengambil list data output dari transaksi lain yang dimiliki oleh user yang masih belum dipakai
	unspent := GetUnspent(myaddress)
	// mengurutkan list data output dari yang memiliki value terbesar hingga yang terkecil
	OrderUnspent(&unspent)
	// membuat array dari tipe bentukan "Destination", yang berisi address tujuan dan jumlah yang ingin dikirimkan
	dest := make([]Destination, 0)
	// menambah elemen array dari address tujuan yang telah didefinisikan sebelumnya
	dest = append(dest, Destination{destaddress, sendvalue})
	// menghitung total fee dan jumlah list output yang dipakai untuk pengiriman sejumlah doge yang telah didefinisikan
	// serta menambah elemen array destinasi yang baru ke address user jika ada kembalian doge yang perlu dibayar ke user
	totalfee, numindex := ChangeUnspent(sendvalue, unspent, &dest, myaddress)
	// menampilkan total fee dari proses transaksi
	log.Printf("totalfee     : %v\n", IntToStr(totalfee))
	// mengecek apakah jumlah saldo masih mencukupi, jika iya, proses pembuatan hex transaksi dijalankan
	if StrToInt(balance.Balance) >= (sendvalue + totalfee) {
		// membuat hex transaksi yang sudah di signature
		signtx := CreateSignedTransaction(unspent, dest, wallet, numindex)
		// menampilkan hex transaksi yang sudah di signature di console
		log.Printf("signtxhex    : %v\n", signtx)
		// konversi hex transaksi yang sudah di signature ke dalam tipe byte untuk keperluan broadcasting
		bin, _ := hex.DecodeString(signtx)
		if err := DogeBroadcaster.Broadcast(bin); err != nil {
			log.Print(err)
			t.Fail()
			return
		}
		// jika jumlah saldo tidak mencukupi jumlah pembayaran beserta total fee, menampilkan keterangan di console
	} else if StrToInt(balance.Balance) >= sendvalue {
		log.Printf("total fee belum melewati batas minimum, transaksi tidak dapat dilakukan\n")
		// jika jumlah saldo lebih kecil dari jumlah pembayaran, menampilkan keterangan di console
	} else {
		log.Printf("saldo tidak mencukupi, transaksi tidak dapat dilakukan\n")
	}
}
