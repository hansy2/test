package main

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
	"os"

	"log"
	"os/exec"
	"time"
)

func printOK() {
	fmt.Println("OK")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkedResult(rows []mysql.Row, res mysql.Result, err error) ([]mysql.Row,
	mysql.Result) {
	checkError(err)
	return rows, res
}

func main() {

	for {
		f
		time.Sleep(5000 * time.Millisecond)
		neu()
		time.Sleep(2000 * time.Millisecond)
		rebuild()
		time.Sleep(2000 * time.Millisecond)
		kill()
		time.Sleep(2000 * time.Millisecond)
		start()
		time.Sleep(2000 * time.Millisecond)
		stop()
		time.Sleep(2000 * time.Millisecond)
		neustart()
	}
}

func neu() {

	user := "user"
	pass := "passwort"
	dbname := "shoutcast"
	proto := "tcp"
	addr := "server:3306"
	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Printf("Connect to %s:%s... ", proto, addr)
	checkError(db.Connect())
	printOK()

	fmt.Println("Select from A... ")
	rows, res := checkedResult(db.Query("select * from Servers where action = 'neu' AND server = '1'"))
	id := res.Map("id")
	sisid := res.Map("sisid")
	for ii, row := range rows {

		fmt.Print(row.Int(sisid))

		err5 := exec.Command("mkdir", "/home/"+row.Str(sisid)+"").Run()

		if err5 != nil {
			log.Print(err5)
		}

		err := exec.Command("cp", "/streams/ftp/sc_serv", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err != nil {
			log.Fatal(err)
		}
		err1 := exec.Command("cp", "/streams/ftp/sc_trans", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err1 != nil {
			log.Fatal(err)
		}
		err2 := exec.Command("cp", "/streams/ftp/sc_serv.conf", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err2 != nil {
			log.Fatal(err)
		}
		err3 := exec.Command("cp", "/streams/ftp/sc_trans.conf", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err3 != nil {
			log.Fatal(err)
		}

		checkedResult(db.Query("INSERT INTO `ftpuser` (`id`, `userid`, `passwd`, `uid`, `gid`, `homedir`, `shell`, `count`, `accessed`, `modified`) VALUES (1, 'rico', '708410', 2001, 2001, '/home/rico', '/sbin/nologin', 4, '2012-11-13 13:56:35', '0000-00-00 00:00:00'),"))

		checkedResult(db.Query("UPDATE Servers SET action='no' where sisid ='" + row.Str(sisid) + "' LIMIT 1 "))

		fmt.Printf(
			"\n Row: %d\n id:  %-10s \n sisid: %-8d  \n", ii,
			"'"+row.Str(id)+"'",
			row.Int(sisid),
		)
	}

	printOK()

	fmt.Print("Close connection... ")
	checkError(db.Close())
	printOK()

}

func kill() {

	user := "user"
	pass := "passwort"
	dbname := "shoutcast"
	proto := "tcp"
	addr := "server:3306"
	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Printf("Connect to %s:%s... ", proto, addr)
	checkError(db.Connect())
	printOK()

	fmt.Println("Select from A... ")
	rows, res := checkedResult(db.Query("select * from Servers where action = 'kill' AND server = '1'"))
	id := res.Map("id")
	sisid := res.Map("sisid")
	for ii, row := range rows {

		fmt.Print(row.Int(sisid))

		//exec.Command("mkdir", "/home/"+row.Str(sisid)+"").Run()
		//exec.Command("cp", "/streams/ftp/* -R /home/10002").Run()

		err4 := exec.Command("/bin/sh", "-c", "kill $(ps aux | grep /home/"+row.Str(sisid)+"  | awk '{print $2}')").Run()

		if err4 != nil {
			log.Print(err4)
		}
		err5 := exec.Command("/bin/sh", "-c", "rm -r /home/"+row.Str(sisid)+"").Run()

		if err5 != nil {
			log.Print(err5)
		}

		checkedResult(db.Query("UPDATE Servers SET action='no' where sisid ='" + row.Str(sisid) + "' LIMIT 1 "))

		fmt.Printf(
			"\n Row: %d\n id:  %-10s \n sisid: %-8d  \n", ii,
			"'"+row.Str(id)+"'",
			row.Int(sisid),
		)
	}

	printOK()

	fmt.Print("Close connection... ")
	checkError(db.Close())
	printOK()

}

func start() {

	user := "user"
	pass := "passwort"
	dbname := "shoutcast"
	proto := "tcp"
	addr := "server:3306"
	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Printf("Connect to %s:%s... ", proto, addr)
	checkError(db.Connect())
	printOK()

	fmt.Println("Select from A... ")
	rows, res := checkedResult(db.Query("select * from Servers where action = 'start' AND server = '1'"))
	id := res.Map("id")
	sisid := res.Map("sisid")
	for ii, row := range rows {

		fmt.Print(row.Int(sisid))

		err := exec.Command("/bin/sh", "-c", "nohup /home/"+row.Str(sisid)+"/sc_serv /home/"+row.Str(sisid)+"/sc_serv.conf > /dev/null 2>&1 &").Run()

		if err != nil {
			log.Print(err)
		}

		checkedResult(db.Query("UPDATE Servers SET action='no' where sisid ='" + row.Str(sisid) + "' LIMIT 1 "))

		fmt.Printf(
			"\n Row: %d\n id:  %-10s \n sisid: %-8d  \n", ii,
			"'"+row.Str(id)+"'",
			row.Int(sisid),
		)
	}

	printOK()

	fmt.Print("Close connection... ")
	checkError(db.Close())
	printOK()

}

func stop() {

	user := "user"
	pass := "passwort"
	dbname := "shoutcast"
	proto := "tcp"
	addr := "server:3306"
	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Printf("Connect to %s:%s... ", proto, addr)
	checkError(db.Connect())
	printOK()

	fmt.Println("Select from A... ")
	rows, res := checkedResult(db.Query("select * from Servers where action = 'stop' AND server = '1'"))
	id := res.Map("id")
	sisid := res.Map("sisid")
	for ii, row := range rows {

		fmt.Print(row.Int(sisid))

		err := exec.Command("/bin/sh", "-c", "kill $(ps aux | grep /home/"+row.Str(sisid)+"  | awk '{print $2}')").Run()

		if err != nil {
			log.Print(err)
		}

		checkedResult(db.Query("UPDATE Servers SET action='no' where sisid ='" + row.Str(sisid) + "' LIMIT 1 "))

		fmt.Printf(
			"\n Row: %d\n id:  %-10s \n sisid: %-8d  \n", ii,
			"'"+row.Str(id)+"'",
			row.Int(sisid),
		)
	}

	printOK()

	fmt.Print("Close connection... ")
	checkError(db.Close())
	printOK()

}

func neustart() {

	user := "user"
	pass := "passwort"
	dbname := "shoutcast"
	proto := "tcp"
	addr := "server:3306"
	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Printf("Connect to %s:%s... ", proto, addr)
	checkError(db.Connect())
	printOK()

	fmt.Println("Select from A... ")
	rows, res := checkedResult(db.Query("select * from Servers where action = 'neustart' AND server = '1'"))
	id := res.Map("id")
	sisid := res.Map("sisid")
	for ii, row := range rows {

		fmt.Print(row.Int(sisid))

		err := exec.Command("/bin/sh", "-c", "kill $(ps aux | grep /home/"+row.Str(sisid)+"  | awk '{print $2}')").Run()

		if err != nil {
			log.Print(err)
		}

		err2 := exec.Command("/bin/sh", "-c", "nohup /home/"+row.Str(sisid)+"/sc_serv /home/"+row.Str(sisid)+"/sc_serv.conf > /dev/null 2>&1 &").Run()

		if err2 != nil {
			log.Print(err2)
		}

		checkedResult(db.Query("UPDATE Servers SET action='no' where sisid ='" + row.Str(sisid) + "' LIMIT 1 "))

		fmt.Printf(
			"\n Row: %d\n id:  %-10s \n sisid: %-8d  \n", ii,
			"'"+row.Str(id)+"'",
			row.Int(sisid),
		)
	}

	printOK()

	fmt.Print("Close connection... ")
	checkError(db.Close())
	printOK()

}

func rebuild() {

	user := "user"
	pass := "passwort"
	dbname := "shoutcast"
	proto := "tcp"
	addr := "server:3306"
	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Printf("Connect to %s:%s... ", proto, addr)
	checkError(db.Connect())
	printOK()

	fmt.Println("Select from A... ")
	rows, res := checkedResult(db.Query("select * from Servers where action = 'rebuild' AND server = '1'"))
	id := res.Map("id")
	sisid := res.Map("sisid")
	for ii, row := range rows {

		fmt.Print(row.Int(sisid))

		err4 := exec.Command("/bin/sh", "-c", "kill $(ps aux | grep /home/"+row.Str(sisid)+"  | awk '{print $2}')").Run()

		if err4 != nil {
			log.Print(err4)
		}
		err5 := exec.Command("/bin/sh", "-c", "rm -r /home/"+row.Str(sisid)+"").Run()

		if err5 != nil {
			log.Print(err5)
		}
		err6 := exec.Command("mkdir", "/home/"+row.Str(sisid)+"").Run()

		if err6 != nil {
			log.Print(err6)
		}
		err := exec.Command("cp", "/streams/ftp/sc_serv", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err != nil {
			log.Fatal(err)
		}
		err1 := exec.Command("cp", "/streams/ftp/sc_trans", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err1 != nil {
			log.Fatal(err)
		}
		err2 := exec.Command("cp", "/streams/ftp/sc_serv.conf", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err2 != nil {
			log.Fatal(err)
		}
		err3 := exec.Command("cp", "/streams/ftp/sc_trans.conf", "-R", "/home/"+row.Str(sisid)+"").Run()

		if err3 != nil {
			log.Fatal(err)
		}

		checkedResult(db.Query("UPDATE Servers SET action='no' where sisid ='" + row.Str(sisid) + "' LIMIT 1 "))

		fmt.Printf(
			"\n Row: %d\n id:  %-10s \n sisid: %-8d  \n", ii,
			"'"+row.Str(id)+"'",
			row.Int(sisid),
		)
	}

	printOK()

	fmt.Print("Close connection... ")
	checkError(db.Close())
	printOK()

}
