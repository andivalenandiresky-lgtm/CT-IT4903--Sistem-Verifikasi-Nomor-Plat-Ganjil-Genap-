# CT-IT4903--Sistem-Verifikasi-Nomor-Plat-Ganjil-Genap-

package main

import "fmt"

func main() {
	var hari string
	var bulan string
	var tahun string

	validTanggal := false

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

		angkaHari := 0
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
			fmt.Println("Tanggal tidak valid")
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

			angkaBulan := 0
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
				fmt.Println("Bulan tidak valid")
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
					fmt.Println("Tahun tidak valid")
				} else {

					angkaTahun := 0
					i = 0
					for i < len(tahun) {
						angkaTahun = angkaTahun*10 + int(tahun[i]-'0')
						i = i + 1
					}
					validTanggal = true

					h := angkaHari
					m := angkaBulan
					y := angkaTahun

					if m < 3 {
						m = m + 12
						y = y - 1
					}

					k := y % 100
					j := y / 100

					hasil := (h + (13*(m+1))/5 + k + k/4 + j/4 + 5*j) % 7

					namaHari := ""

					if hasil == 0 {
						namaHari = "Sabtu"
					}
					if hasil == 1 {
						namaHari = "Minggu"
					}
					if hasil == 2 {
						namaHari = "Senin"
					}
					if hasil == 3 {
						namaHari = "Selasa"
					}
					if hasil == 4 {
						namaHari = "Rabu"
					}
					if hasil == 5 {
						namaHari = "Kamis"
					}
					if hasil == 6 {
						namaHari = "Jumat"
					}

					fmt.Println("Hari:", namaHari)
					if namaHari == "Sabtu" || namaHari == "Minggu" {
						fmt.Println("Tidak berlaku ganjil genap (˶ > ₃ < ˶)")
					} else {
						var jenisKendaraan string
						fmt.Print("Masukkan jenis kendaraan (listrik / nonlistrik): ")
						fmt.Scanln(&jenisKendaraan)

						if jenisKendaraan == "listrik" {
							fmt.Println("Kendaraan listrik bebas melintas (˶˃𐃷˂˶)")
						} else {
							var plat string
							fmt.Print("Masukkan dua digit terakhir plat: ")
							fmt.Scanln(&plat)

							validPlat := true
							i = 0
							for i < len(plat) {
								if plat[i] < '0' || plat[i] > '9' {
									validPlat = false
								}
								i = i + 1
							}

							if validPlat {
								digitPlat := int(plat[len(plat)-1] - '0')
								platGanjil := digitPlat%2 == 1
								tanggalGanjil := angkaHari%2 == 1

								if tanggalGanjil == platGanjil {
									fmt.Println("KENDARAAN BOLEH MELINTAS (>⩊<)")
								} else {
									fmt.Println("KENDARAAN DILARANG MELINTAS (¬_¬) !!!")
								}
							} else {
								fmt.Println("Plat tidak valid")
							}
						}
					}
				}
			}
		}
	}
}


