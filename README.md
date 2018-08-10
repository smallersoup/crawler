# crawler
IDEA go插件下载地址
https://plugins.jetbrains.com/plugin/9568-go

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




java:

 /**
     * 日志压缩
     * @param waitZipFile 要压缩文件名
     * @throws IOException
     */
    private void zipAuditLogFile(String waitZipFile) throws IOException {
        File oldFile = new File(waitZipFile);

        if (!oldFile.exists()) {
            LOGGER.error("zipAuditLogFile name is {} not exist", waitZipFile);
            return;
        }

        //生成zip文件名
        DateFormat dataFormat = new SimpleDateFormat(AUDIT_LOG_FORMAT);
        String formatTime = dataFormat.format(oldFile.lastModified());

        int end = waitZipFile.length() - AUDIT_FILE_EXT.length();
        String zipFileName = waitZipFile.subSequence(0, end) + "_" + formatTime + AUDIT_FILE_ZIP_SUFFIX;

        File zipFile = new File(zipFileName);

        FileOutputStream zipfos = null;
        ZipOutputStream zipOs = null;
        CheckedOutputStream cos = null;


        try {
            zipfos = new FileOutputStream(zipFile);
            cos = new CheckedOutputStream(zipfos, new CRC32());

            zipOs = new ZipOutputStream(cos);

            String baseDir = (this.rootPath != null && this.rootPath.endsWith(File.separator)) ? this.rootPath : (this.rootPath + File.separator);

            compress(oldFile, zipOs, baseDir);

            if (zipFile.exists()) {
                // 写完的日志文件权限改为400
                try
                {
//                    Path path = Paths.get(zipFile.getAbsolutePath());
                    boolean isSuccess = PathUtil.setFilePermision(zipFile.toPath(), ARCHIVE_LOGFILE_PERMISION);
                    LOGGER.warn("set archive file: {}, permision result is {}", zipFile.getAbsolutePath(), isSuccess);
                }
                catch (IOException e)
                {
                    LOGGER.error("set archive file permision catch an error: {}", e);
                }
            }

        } finally {

            if(null != zipOs){
                zipOs.close();
            }

            if(null != cos){
                cos.close();
            }

            if(null != zipfos){
                zipfos.close();
            }
        }
    }

    /**
     * 压缩文件或目录
     * @param oldFile 要压缩的文件
     * @param zipOut 压缩文件流
     * @param baseDir baseDir
     * @throws IOException
     */
    private void compress(File oldFile, ZipOutputStream zipOut, String baseDir) throws IOException {

        if (oldFile.isDirectory()) {

            compressDirectory(oldFile, zipOut, baseDir);

        }else {
            compressFile(oldFile, zipOut, baseDir);
        }
    }

    /**
     * 压缩目录
     * @param dir 要压缩的目录
     * @param zipOut 压缩文件流
     * @param baseDir baseDir
     * @throws IOException
     */
    private void compressDirectory(File dir, ZipOutputStream zipOut, String baseDir) throws IOException {

        File[] files = dir.listFiles();

        for (File file: files) {
            compress(file, zipOut, baseDir + dir.getName() + File.separator);
        }
    }

    /**
     * 压缩文件
     * @param oldFile 要压缩的文件
     * @param zipOut 压缩文件流
     * @param baseDir baseDir
     * @throws IOException
     */
    private void compressFile(File oldFile, ZipOutputStream zipOut, String baseDir) throws IOException {

        if (!oldFile.exists()) {
            LOGGER.error("zipAuditLogFile name is {} not exist", oldFile);
            return;
        }

        BufferedInputStream bis = null;

        try {

            bis = new BufferedInputStream(new FileInputStream(oldFile));

            ZipEntry zipEntry = new ZipEntry(baseDir + oldFile.getName());

            zipOut.putNextEntry(zipEntry);

            int count;

            byte data[] = new byte[ZIP_BUFFER];

            while ( (count = bis.read(data, 0, ZIP_BUFFER)) != -1 ) {
                zipOut.write(data, 0, count);
            }

        } finally {
            if (null != bis) {
                bis.close();
            }
        }

    }
