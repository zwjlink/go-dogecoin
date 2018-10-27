package dogecoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// mengecek dan memperbaiki value r atau s dari signature agar tidak bernilai negatif
func SignCorrect(sign string) string {
	/*seperti yang diketahui, panjang r ataupun s yang dihasilkan oleh signature tidak selamanya pas 32 byte atau 64 karakter
	oleh karenanya pada bagian di bawah ini, jika seandainya panjang r atau s kurang dari 32 byte maka pada r atau s
	ditambahkan prefix angka nol hingga panjangnya jadi pas 32 byte*/
	for len(sign) < 64 {
		sign = "0" + sign
	}
	/*kalau pada bagian ini, byte prefix diambil dari r ataupun s*/
	prefix, err := strconv.ParseUint(sign[:2], 16, 64)
	ErrorCheck(err)
	/*selanjutnya dicek apakah nilainya lebih dari 0x7f. Jika ya, berarti nilai r atau s yang diperoleh bertanda negatif
	untuk menghindari nilai negatif tersebut, maka ditambahkan byte prefix sebesar nol "00" agar nilainya jadi tidak negatif*/
	if prefix >= 0x7f {
		sign = "00" + sign
	}
	/*mengembalikan nilai r atau s yang sudah diperbaiki agar tidak negatif*/
	return sign
}

// memperbaiki format hex agar panjangnya bernilai genap
func EvenCorrect(num int) string {
	var numstring bytes.Buffer
	/*pada bagian ini, bilangan integer num dituliskan dalam format hex, kemudian panjangnya dalam format hex dicek
	jika panjangnya tidak genap, maka ditambahkan angka nol di depannya agar jadi genap panjangnnya*/
	if (len(fmt.Sprintf("%x", num)) % 2) > 0 {
		numstring.WriteString("0" + fmt.Sprintf("%x", num))
		/*kalau panjangnnya udah genap, tinggal tulis bilangan integer num tadi dalam format hex*/
	} else {
		numstring.WriteString(fmt.Sprintf("%x", num))
	}
	/*mengembalikan bilangan integer num dalam format hex dengan panjang yang sudah diperbaiki agar genap*/
	return numstring.String()
}

// sama seperti fungsi evencorrect, dengan penambahan prefix jika syarat batas value tertentu telah dilewati
func VarInt(num int) string {
	var numstring, numfinal bytes.Buffer
	/*memperbaiki bilangan integer num dalam format hex agar panjangnya jadi genap dengan memanggil fungsi EvenCorrect*/
	numstring.WriteString(EvenCorrect(num))
	/*aturan dari dokumentasi, kalau nilai integer num melewati batas-batas yang telah ditentukan, maka ditambahkan prefix tertentu*/
	switch {
	case num <= 0xfc:
		//do nothing
	case num <= 0xffff:
		numfinal.WriteString("fd")
	case num <= 0xffffffff:
		numfinal.WriteString("fe")
	default:
		numfinal.WriteString("ff")
	}
	/*karena format penulisannya harus dalam little endian, maka urutan penulisannya perlu dibalikkan dengan fungsi ReverseHex*/
	numfinal.WriteString(ReverseHex(numstring.String()))
	/*mengembalikan integer num dalam format hex dan little endian, serta sesuai denga aturan dokumentasi*/
	return numfinal.String()
}

// mengkonversi value dari doge dalam bentuk string ke dalam bentuk unsigned integer-nya
func StrToInt(doge_value string) uint64 {
	var value uint64
	var err error
	/*value dari doge dalam string, tanda titiknya dihilangkan agar dapat dikonversi ke integer*/
	doge_string_value := fmt.Sprint(strings.Replace(doge_value, ".", "", -1))
	/*kalau tidak ada value yang tercantum, diasumsikan bernilai nol*/
	if doge_string_value == "" {
		value = 0
		/*kalau valuenya tidak kosong, proses konversi dilakukan*/
	} else {
		value, err = strconv.ParseUint(doge_string_value, 10, 64)
		ErrorCheck(err)
	}
	/*mengembalikan value yang sudah dikonversi*/
	return value
}

// konversi value unsigned integer dari doge ke dalam bentuk string-nya
func IntToStr(doge_value uint64) string {
	return fmt.Sprintf("%v.%08v", doge_value/100000000, doge_value%100000000)
}

// memperbaiki bagian hex tertentu dalam hex transaksi agar berada dalam format little endian
func ReverseHex(hexa string) string {
	/*konversi string hexadecimal dalam format byte*/
	bytes, err := hex.DecodeString(hexa)
	ErrorCheck(err)
	var temp byte
	/*proses pembalikan urutan byte dilakukan sepanjang loop ini*/
	for i := 0; i < (len(bytes) - (i + 1)); i++ {
		temp = bytes[i]
		bytes[i] = bytes[len(bytes)-(i+1)]
		bytes[len(bytes)-(i+1)] = temp
	}
	/*konversi byte yang sudah dibalikkan urutannya tadi ke dalam string hexadecimal*/
	reversed := hex.EncodeToString(bytes)
	/*mengembalikan string hexadecimal yang sudah dibalikkan*/
	return reversed
}

// melakukan hash dengan algoritma sha256 sebanyak dua kali
func Hash(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:]
}

// mengecek dan menampilkan error di console
func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
