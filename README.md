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

/**
 * 提供文件路径公共函数 改变权限，判断是否正规文件，判断是否路径在安全路径下等
 */
import java.io.IOException;
import java.nio.file.FileSystems;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.attribute.PosixFilePermission;
import java.nio.file.attribute.PosixFilePermissions;
import java.util.Set;

/**
 * 路径工具
 * 
 */
public class PathUtil implements IPathUtil
{
    private static final boolean ISPOSIX = FileSystems.getDefault().supportedFileAttributeViews().contains("posix");

    private static final String FILE_PERMISION = "600";

    /**
     * 设置默认的文件权限
     * 
     * @param file
     *            文件路径
     * @return true 设置成功
     */
    public static boolean setDefaultPermision(Path file) throws IOException
    {
        return setFilePermision(file, FILE_PERMISION);
    }

    /**
     * 设置文件的权限，尽在posix下有效
     * 
     * @param file
     *            文件
     * @param perm
     *            权限 类似 “rw-r-----”, "750"
     * @return true 修改成功 false 修改失败
     * @throws IOException
     */
    public static boolean setFilePermision(Path file, String perm) throws IOException
    {
        if (!ISPOSIX)
        {
            return true;
        }
        // 750 -> "rwxr-x---"
        if (perm.length() == 3)
        {
            perm = trans2StrPerm(perm);
        }

        if (perm.length() != 9)
        {
            return false;
        }
        Set<PosixFilePermission> perms = PosixFilePermissions.fromString(perm);
        Files.setPosixFilePermissions(file, perms);
        return true;
    }

    /**
     * 转换
     * 
     * @param digitPerm
     *            长度为3的数字字符串
     * @return
     */
    private static String trans2StrPerm(String digitPerm)
    {
        StringBuilder builder = new StringBuilder(9);
        builder.append(toStringPerm(digitPerm.charAt(0)));// owner
        builder.append(toStringPerm(digitPerm.charAt(1)));// group
        builder.append(toStringPerm(digitPerm.charAt(2)));// other
        return builder.toString();
    }

    private static String toStringPerm(char ch)
    {
        switch (ch - '0')
        {
            case 7:
                return "rwx";
            case 6:
                return "rw-";
            case 5:
                return "r-x";
            case 4:
                return "r--";
            case 3:
                return "-wx";
            case 2:
                return "-w-";
            case 1:
                return "--x";
            case 0:
                return "---";
            default:
                return "";
        }
    }

}

  ```
