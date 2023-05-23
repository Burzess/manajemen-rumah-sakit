package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strukturdataMVC/view"
	dokter "strukturdataMVC/view/dokterview"
	pasien "strukturdataMVC/view/pasienview"
	suster "strukturdataMVC/view/susterview"
)

func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	var menu int
	for menu != 4 {
		Clear()
		view.MenuUtama()
		fmt.Scan(&menu)
		switch menu {
		case 1:
			var pilih int
			var back string
			for pilih != 6 {
				Clear()
				view.SubMenu()
				fmt.Scan(&pilih)
				switch pilih {
				case 1:
					Clear()
					dokter.InsertDokter()
					fmt.Print("kembali (y/t): ")
					fmt.Scan(&back)
					if back == "t" || back == "T" {
						os.Exit(0)
					}
				case 2:
					Clear()
					dokter.UpdateDoter()
					fmt.Print("kembali (y/t): ")
					fmt.Scan(&back)
					if back == "t" || back == "T" {
						os.Exit(0)
					}
				case 3:
					Clear()
					dokter.DeleteDoter()
					fmt.Print("kembali (y/t): ")
					fmt.Scan(&back)
					if back == "t" || back == "T" {
						os.Exit(0)
					}
				case 4:
					Clear()
					dokter.ViewById()
					fmt.Print("kembali (y/t): ")
					fmt.Scan(&back)
					if back == "t" || back == "T" {
						os.Exit(0)
					}
				case 5:
					Clear()
					dokter.ViewDokter()
					fmt.Print("kembali (y/t): ")
					fmt.Scan(&back)
					if back == "t" || back == "T" {
						os.Exit(0)
					}
				case 6:
				default:
					Clear()
					println("pilihan tidak valid")
				}
			}
		case 2:
			Clear()
			view.MenuUtama()
			fmt.Scan(&menu)
			switch menu {
			case 1:
				var pilih int
				var back string
				for pilih != 6 {
					Clear()
					view.SubMenu()
					fmt.Scan(&pilih)
					switch pilih {
					case 1:
						Clear()
						suster.InsertSuster()
						fmt.Print("kembali (y/t): ")
						fmt.Scan(&back)
						if back == "t" || back == "T" {
							os.Exit(0)
						}
					case 2:
						Clear()
						suster.UpdateDoter()
						fmt.Print("kembali (y/t): ")
						fmt.Scan(&back)
						if back == "t" || back == "T" {
							os.Exit(0)
						}
					case 3:
						Clear()
						suster.DeleteDoter()
						fmt.Print("kembali (y/t): ")
						fmt.Scan(&back)
						if back == "t" || back == "T" {
							os.Exit(0)
						}
					case 4:
						Clear()
						suster.ViewById()
						fmt.Print("kembali (y/t): ")
						fmt.Scan(&back)
						if back == "t" || back == "T" {
							os.Exit(0)
						}
					case 5:
						Clear()
						suster.ViewSuster()
						fmt.Print("kembali (y/t): ")
						fmt.Scan(&back)
						if back == "t" || back == "T" {
							os.Exit(0)
						}
					case 6:
					default:
						Clear()
						println("pilihan tidak valid")
					}
				}
			case 3:
				Clear()
				view.MenuUtama()
				fmt.Scan(&menu)
				switch menu {
				case 1:
					var pilih int
					var back string
					for pilih != 6 {
						Clear()
						view.SubMenu()
						fmt.Scan(&pilih)
						switch pilih {
						case 1:
							Clear()
							pasien.InsertPasien()
							fmt.Print("kembali (y/t): ")
							fmt.Scan(&back)
							if back == "t" || back == "T" {
								os.Exit(0)
							}
						case 2:
							Clear()
							pasien.UpdateDoter()
							fmt.Print("kembali (y/t): ")
							fmt.Scan(&back)
							if back == "t" || back == "T" {
								os.Exit(0)
							}
						case 3:
							Clear()
							pasien.DeleteDoter()
							fmt.Print("kembali (y/t): ")
							fmt.Scan(&back)
							if back == "t" || back == "T" {
								os.Exit(0)
							}
						case 4:
							Clear()
							pasien.ViewById()
							fmt.Print("kembali (y/t): ")
							fmt.Scan(&back)
							if back == "t" || back == "T" {
								os.Exit(0)
							}
						case 5:
							Clear()
							pasien.ViewPasien()
							fmt.Print("kembali (y/t): ")
							fmt.Scan(&back)
							if back == "t" || back == "T" {
								os.Exit(0)
							}
						case 6:
						default:
							Clear()
							println("pilihan tidak valid")
						}
					}
				case 4:
					Clear()
					fmt.Println("Terimakasih")
				default:
					Clear()
					println("pilih tidak valid")
				}
			}
		}
	}
}
