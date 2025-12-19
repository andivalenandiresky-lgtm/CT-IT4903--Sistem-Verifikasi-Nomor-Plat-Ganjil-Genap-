# CT-IT4903--Sistem-Verifikasi-Nomor-Plat-Ganjil-Genap-

package main

import "fmt"

func main() {

	fmt.Println(" >⩊< SELAMAT DATANG DI PROGRAM SISTEM VERIFIKASI PLAT >⩊<")

	ulangProgram := true

	for ulangProgram {

		var hari string
		var bulan string
		var tahun string

		validTanggal := false
		angkaHari := 0
		angkaBulan := 0
		angkaTahun := 0

		for validTanggal == false {

			fmt.Print("Masukkan tanggal (1-31): ")
			fmt.Scanln(&hari)

			validHari := true
			i := 0

			if len(hari) == 0 {
				validHari = false
			}

			for i < len(hari) {
				if hari[i] < '0' || hari[i] > '9' {
					validHari = false
				}
				i = i + 1
			}

			angkaHari = 0
			if validHari {
				i = 0
				for i < len(hari) {
					angkaHari = angkaHari*10 + int(hari[i]-'0')
					i = i + 1
				}
				if angkaHari < 1 || angkaHari > 31 {
					validHari = false
				}
			}

			if validHari == false {
				fmt.Println("Tanggal tidak valid ( •̀ ᴖ •́ )")
			} else {
				
				fmt.Print("Masukkan bulan (1-12): ")
				fmt.Scanln(&bulan)

				validBulan := true
				i = 0

				for i < len(bulan) {
					if bulan[i] < '0' || bulan[i] > '9' {
						validBulan = false
					}
					i = i + 1
				}

				angkaBulan = 0
				if validBulan {
					i = 0
					for i < len(bulan) {
						angkaBulan = angkaBulan*10 + int(bulan[i]-'0')
						i = i + 1
					}
					if angkaBulan < 1 || angkaBulan > 12 {
						validBulan = false
					}
				}

				if validBulan == false {
					fmt.Println("Bulan tidak valid ( •̀ ᴖ •́ )")
				} else {

					fmt.Print("Masukkan tahun: ")
					fmt.Scanln(&tahun)

					validTahun := true
					i = 0

					for i < len(tahun) {
						if tahun[i] < '0' || tahun[i] > '9' {
							validTahun = false
						}
						i = i + 1
					}

					if validTahun == false {
						fmt.Println("Tahun tidak valid ( •̀ ᴖ •́ )")
					} else {

						angkaTahun = 0
						i = 0
						for i < len(tahun) {
							angkaTahun = angkaTahun*10 + int(tahun[i]-'0')
							i = i + 1
						}

						validTanggal = true
					}
				}
			}
		}
