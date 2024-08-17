package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	dictFileName       = "./dict.txt"
	dictBackupFileName = "./dict_bak_%d.txt"
)

var (
	logger      = slog.New(slog.NewTextHandler(os.Stderr, nil))
	nameDictMap = map[string]struct{}{}
)

func init() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[当前工作目录：%s]\n", dir)

	// 读取dict文件
	// 打开文件
	file, err := os.Open(dictFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保在函数结束时关闭文件

	backupFileName := fmt.Sprintf(dictBackupFileName, time.Now().Unix())
	fmt.Println("[开始备份文件] 文件名: ", backupFileName)

	backupFile, err := os.Create(backupFileName)
	if err != nil {
		log.Fatal("备份字典文件错误，err: ", err.Error())
	}
	defer backupFile.Close()

	// 创建 Scanner 来按行读取
	scanner := bufio.NewScanner(file)

	// 使用 Scan() 方法按行迭代文件
	for scanner.Scan() {
		line := scanner.Text() // 获取当前行的文本
		// 加载到内存中的Map
		nameDictMap[line] = struct{}{}

		// 写入备份文件
		_, err = backupFile.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("写入备份文件时发生错误: %v", err)
			continue
		}
	}
	// 检查是否有可能的错误
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("[备份完成!!!] 文件名: ", backupFileName)
}

func main() {
	inputArr := scanInputText()
	tempDict := buildDict(inputArr)
	saveDict2Txt(tempDict)
}

func saveDict2Txt(tempDict []string) {
	for _, v := range tempDict {
		// 将新的字典数据覆盖之前的数据
		nameDictMap[v] = struct{}{}
	}

	// 创建文件，如果文件已存在，它将被截断（覆盖）
	file, err := os.Create(dictFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保在函数结束时关闭文件

	var fileLines []string
	for v, _ := range nameDictMap {
		fileLines = append(fileLines, v)
	}
	fileText := strings.Join(fileLines, "\n")

	// 写入数据到文件
	_, err = file.WriteString(fileText)
	if err != nil {
		logger.Error("覆盖写入 dict.txt 错误 , err: ", err.Error())
		return
	}
}

// 对预处理文本进行清理
func strTrims(str string) string {
	// 去两边空格
	result := strings.TrimSpace(str)
	// 清理额外字符
	cutSets := []string{"(", ")", " ", "!", "！", "单价"} // ","

	for _, cut := range cutSets {
		result = strings.Trim(result, cut)
	}

	return result
}

// 获取数字出现的第一个位置
func getNumIndex(str string) int {
	// 正则表达式匹配字符串中的第一个数字
	re := regexp.MustCompile(`\d+`)
	matches := re.FindStringIndex(str)

	if len(matches) == 2 {
		// matches[0] 是匹配的数字的起始索引，matches[1] 是结束索引
		return matches[0]
	}
	return -1
}

// 分词按行输出字典数据
func buildDict(textArr []string) []string {
	results := []string{}
	// 按行处理识别的字符数据
	for i, _ := range textArr {
		curText := textArr[i]
		// 获取该行文本汇总的数字出现的位置
		numIndex := getNumIndex(curText)
		// 没有数据则本行直接保留
		if numIndex == -1 {
			// 清洗数据
			tempText := strTrims(curText)
			if len(tempText) > 0 {
				results = append(results, tempText)
			}
			continue
		} else {
			str1 := strTrims(curText[0:numIndex])
			str2 := strTrims(curText[numIndex:])

			if len(str1) > 0 {
				results = append(results, str1)
			}

			// 价格数字
			if len(str2) > 0 {
				index := getNumIndex(str2)
				if index == 0 {
					fmt.Println("这是价格：", str2)
					continue
				}
			}
		}
	}

	fmt.Println("=============================================")
	for _, v := range results {
		fmt.Println(v)
	}

	return results
}

// 按行接收输入的多行数据 直到回车结束
func scanInputText() []string {
	// 创建一个bufio.Scanner，用于读取控制台输入
	scanner := bufio.NewScanner(os.Stdin)

	// 打印提示信息
	fmt.Println("请输入多行数据，输入空行结束：")
	var inputTextArr []string
	// 使用循环读取每一行输入
	for scanner.Scan() {
		// 读取的文本赋值给text变量
		text := scanner.Text()
		inputTextArr = append(inputTextArr, text)
		// 检查是否输入了空行
		if text == "" {
			break
		}
	}

	// 检查是否有可能发生的错误
	if err := scanner.Err(); err != nil {
		logger.Error("读取输入时发生错误: ", err.Error())
		return nil
	}

	return inputTextArr
}
