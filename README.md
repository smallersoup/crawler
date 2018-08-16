# crawler
IDEA go插件下载地址
https://plugins.jetbrains.com/plugin/9568-go

```go

func startCheckDeleteLogTimer(file string) {

	//get now time
	now := time.Now()

	//get duration of now and 00:00:00 as Timer duration
	du := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Sub(now)

	//count down timer
	t := time.NewTimer(du)

	isFirst := make(chan bool, 1)
	select {
	case isFirstTime := <-t.C:
		fmt.Println("first time for exec delete auditlog check.. FirstTime: ", isFirstTime)
		t.Stop()
		checkDeleteAuditLog(file)
		isFirst <- false

	}

	//this goroutine block until count down timer stop
	<-isFirst

	//ticker for exec delete auditlog check every 24 hours
	checkDelTicker := time.Tick(24 * time.Hour)

	for {
		select {
		case <- checkDelTicker:
			fmt.Println("will exec delete auditlog check.. CurrentTime: ", time.Now())
			checkDeleteAuditLog(file)
		}
	}
}


err := zipFile(config.LogFile, backfile)
	if err != nil {
		fmt.Println(fmt.Sprintf("zip file %s to %s error : %v", config.LogFile, backfile, err))
		return
	} else {
		os.Remove(config.LogFile)
	}

	al.createFile()

	os.Chmod(backfile, 0400)
	al.deleteOldBackfiles(dir)
  
  
  
  func zipFile(source, target string) error {

	zipFile, err := os.OpenFile(target, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0440)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			header.Method = zip.Deflate
		}
		header.SetModTime(time.Now().UTC())
		header.Name = path
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)

		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})
}
```



java:
```java

  ```
