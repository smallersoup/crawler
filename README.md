# crawler
IDEA go插件下载地址
https://plugins.jetbrains.com/plugin/9568-go

```go
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
