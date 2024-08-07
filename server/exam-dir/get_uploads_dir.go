package examdir

import (
	"log"
	"os"
	"path/filepath"
)

// 判断文件夹是否存在
// 这行是一个注释，用于说明这个函数的作用，即检查给定的路径是否存在。
// 需要注意的是，这个函数实际上检查的是路径是否存在，并不特别区分是文件还是文件夹（目录）。

func PathExists(path string) (bool, error) {
	// 定义了一个名为PathExists的函数，它接收一个字符串参数path（表示要检查的路径），
	// 并返回两个值：一个布尔值表示路径是否存在，和一个error值表示可能发生的错误。

	_, err := os.Stat(path)
	// 使用os包的Stat函数尝试获取给定路径的文件信息。
	// os.Stat返回一个FileInfo接口和一个error值。这里我们忽略FileInfo接口的返回值（用_表示），
	// 只关心是否有错误发生。如果路径存在，err将为nil；如果路径不存在或发生其他错误，err将包含错误信息。

	if err == nil {
		// 如果err为nil，表示没有错误发生，也就是说路径存在。
		// 因此，函数返回true（表示路径存在）和nil（表示没有错误）。
		return true, nil
	}

	if os.IsNotExist(err) {
		// 如果err不为nil，我们进一步检查这个错误是否表示“路径不存在”。
		// os.IsNotExist是一个函数，用于检查给定的error值是否表示“文件或目录不存在”的错误。
		// 如果是，函数返回false（表示路径不存在）和nil（因为没有发生除“路径不存在”之外的其他错误）。
		return false, nil
	}

	return false, err
	// 如果err既不为nil，也不表示“路径不存在”，那么说明发生了其他类型的错误（如权限不足等）。
	// 在这种情况下，函数返回false（因为路径不存在或无法确定其存在性）和err（表示发生的具体错误）。
}
func GetUpLoadsDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	exPath := filepath.Dir(ex)
	// 在当前目录下创建一个uploads目录
	fileLoads := filepath.Join(exPath, "uploads")
	// 如果不存在就创建uploads目录
	ok, err := PathExists(fileLoads)
	if err != nil {
		log.Fatalln(err)
	}
	// 不存在目录就创建目录
	if !ok {
		err = os.Mkdir(fileLoads, os.ModePerm) // 权限为 可读可写可执行
		if err != nil {
			log.Fatalln(err)
		}
	}
	return fileLoads
}
