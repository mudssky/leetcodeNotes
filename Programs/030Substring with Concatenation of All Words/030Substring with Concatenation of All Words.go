package main

import (
	"fmt"
	"strings"
)

// 解法一，穷举全排列 求出所有字符串的全排列存到map里，再遍历字符串，用切片取出给定长度的字符串，查map里有没有，如果有，把下标加到结果数组，
// 最后返回结果数组
// 时间复杂度 主要的运算在穷举全排列上，这个部分随时间增长，n!级扩大，故时间复杂度O(n!)
// 空间复杂度 存储n!中排列需要消耗大量空间   空间复杂度O(n!)
// 无法通过测试，并且在自己电脑上运行的结果，内存被占满，算了半天都算不出来

// 递归方式求给定字符串数组的所有排列方式

// a1+a2+a3+a4,第一个位置可以是a1，a2，a3，a4...，分类讨论，然后就变成求 a1 + (a2+a3+...) a2+(a1,a3,a4...)
// 这样就可以用递归的思路， 每次对一个位置分类讨论，运用交换位置法
func getWordsDict(words *[]string) map[string]bool {
	dict := make(map[string]bool)
	var permutation func(dict map[string]bool, words *[]string, index int)
	permutation = func(dict map[string]bool, words *[]string, index int) {
		// 如果确定位置已经到数组的末尾，就把此时这个数组拼接加入map
		// 这是退出条件
		if index == len(*words) {
			dict[strings.Join(*words, "")] = true
		} else {
			// 遍历words数组，使index位置分别和其他位置交换
			for i := index; i < len(*words); i++ {
				(*words)[index], (*words)[i] = (*words)[i], (*words)[index]
				// 确定index的字符，继续往下确定其他字符的位置
				permutation(dict, words, index+1)
				// 得到结果后，把位置交换回来
				(*words)[index], (*words)[i] = (*words)[i], (*words)[index]
			}
		}
	}
	permutation(dict, words, 0)
	return dict
}

/*
func findSubstring(s string, words []string) []int {
	// 求拼接后字符串的长度,用于for循环,如果剩下的字符串长度小于这个拼接字符串长度,那么就不用比较下去了.
	wordsLen := len(strings.Join(words, ""))
	// 如果给出的字符串数组为空,那么返回nil
	if wordsLen == 0 {
		return nil
	}
	sLen := len(s)
	dict := getWordsDict(&words)
	res := []int{}
	// 遍历字符串,切片和拼接字符串相同大小进行比较,如果匹配,首字母下标加入结果数组
	for i := 0; i <= sLen-wordsLen; i++ {
		tmp := s[i : i+wordsLen]
		if _, ok := dict[tmp]; ok {
			res = append(res, i)
		}
	}
	return res
}*/
// 解法二 将字符串数组全部放入map，重复的个数+1表示，然后遍历字符串匹配。
// s切片一个单词长度的进行匹配，如果map能找到，那么每个单词长度为步长遍历一个单词的长度，用一个匹配指数来记录，每匹配一个，匹配指数+1，同时map中的计数器-1
// 如果map中的不够-1了，说明匹配失败，匹配失败就break。结束后，判断匹配指数是否和单词个数相同，相同就说明全匹配上了，然后重置map。继续字符串向下搜索
// 时间复杂度，O(m*n)
// 空间复杂度，主要用到一个map O(n)
// 这个过程中word重复匹配了很多次，显然，还有继续优化的空间
/*
Runtime: 188 ms, faster than 18.35% of Go online submissions for Substring with Concatenation of All Words.
Memory Usage: 7.5 MB, less than 25.00% of Go online submissions for Substring with Concatenation of All Words.
*/
/*func findSubstring(s string, words []string) []int {
	wordsLen := len(words)
	// 如果字符串为空，不用进行匹配，返回nil
	if wordsLen == 0 {
		return nil
	}
	// 每个字符串的长度
	wordLen := len(words[0])

	// 把所有单词加入map用于检索和计算
	wordsDict := make(map[string]int)
	// 把字符串都放到map里，如果有重复的，用value值表示个数
	for _, v := range words {
		if _, ok := wordsDict[v]; ok {
			wordsDict[v]++
		} else {
			wordsDict[v] = 1
		}
	}
	// 拼接后字符串的长度
	wordsJoinStrLen := wordLen * wordsLen
	res := []int{}
	sLen := len(s)
	// 遍历字符串，遍历开头的字母
	for i := 0; i <= sLen-wordsJoinStrLen; i++ {
		// 用一个tmp的map在循环中用于查找和计算，循环结束以后重置一遍为初始数值
		tmpWordsDict := make(map[string]int)
		for inx, val := range wordsDict {
			tmpWordsDict[inx] = val
		}
		// 匹配的个数，当等于所有字符个数，说明所有字符都匹配成功了
		matchedWord := 0
		// 第一个word匹配成功，继续向后匹配
		if _, ok := tmpWordsDict[s[i:i+wordLen]]; ok {
			tmpWordsDict[s[i:i+wordLen]]--
			matchedWord++
			for j := i + wordLen; j <= i+wordsJoinStrLen-wordLen; j += wordLen {
				if newCount, ook := tmpWordsDict[s[j:j+wordLen]]; ook && newCount > 0 {
					tmpWordsDict[s[j:j+wordLen]]--
					matchedWord++
				} else {
					// 如果有任意一个单词没匹配上，break
					break
				}
			}
			// 如果最后匹配的数目达到字符串总数，说明匹配成功
			if matchedWord == wordsLen {
				res = append(res, i)
			}
		}
	}
	return res
}*/

//  解法四，和我的解法接近，算是基础上的优化吧，最大的区别是，是以所有单词长度和为窗口，从后往前进行比较的
// 并且当匹配失败的时候，会把匹配的位置加到之前最后一次成功的位置，也就是说，完美利用了每一次成功的匹配
// 此外，没有map拷贝的的过程，这一步就减少了消耗

func findSubstring(s string, words []string) []int {
	result := []int{}

	if len(words) == 0 {
		return result
	}

	wordMap := map[string]int{}
	for _, v := range words {
		wordMap[v] = wordMap[v] + 1
	}
	wordLen := len(words[0])
	window := len(words) * wordLen
	// i是遍历的粒度，如果以window的长度遍历，初始0-wordLen的部分，单词匹配的状况是不会重复的，之后就会出现重复
	for i := 0; i < wordLen; i++ {
		for j := i; j+window <= len(s); j += wordLen {
			tmpStr := s[j : j+window]
			temMap := make(map[string]int)

			for k := len(words) - 1; k >= 0; k-- {
				word := tmpStr[k*wordLen : (k+1)*wordLen]
				temMap[word]++
				if temMap[word] > wordMap[word] {
					j = j + k*wordLen
					break
				} else if k == 0 {
					//这里不用担心存在tempMap[word] < wordMap[word]的情况，
					//在单词窗口长度与words单词长度和相等的情况下，
					//若存在某个单词wordMap有而tempMap没有，
					//那么必然存在另外一个单词tempMap有而wordMap没有。
					//即必定会进入到上面if的语句中。
					result = append(result, j)
				}
			}
		}
	}
	return result
}

func main() {
	words := []string{"foo", "bar"}
	s := "barfoothefoobarman"
	s2 := "pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbcjjivtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbjjkgcztkcbwagtcnrtqryuqixtzhaakjlurnumzyovawrcjiwabuwretmdamfkxrgqgcdgbrdbnugzecbgyxxdqmisaqcyjkqrntxqmdrczxbebemcblftxplafnyoxqimkhcykwamvdsxjezkpgdpvopddptdfbprjustquhlazkjfluxrzopqdstulybnqvyknrchbphcarknnhhovweaqawdyxsqsqahkepluypwrzjegqtdoxfgzdkydeoxvrfhxusrujnmjzqrrlxglcmkiykldbiasnhrjbjekystzilrwkzhontwmehrfsrzfaqrbbxncphbzuuxeteshyrveamjsfiaharkcqxefghgceeixkdgkuboupxnwhnfigpkwnqdvzlydpidcljmflbccarbiegsmweklwngvygbqpescpeichmfidgsjmkvkofvkuehsmkkbocgejoiqcnafvuokelwuqsgkyoekaroptuvekfvmtxtqshcwsztkrzwrpabqrrhnlerxjojemcxel"
	words2 := []string{"dhvf", "sind", "ffsl", "yekr", "zwzq", "kpeo", "cila", "tfty", "modg", "ztjg", "ybty", "heqg", "cpwo", "gdcj", "lnle", "sefg", "vimw", "bxcb"}
	// fmt.Println(getWordsDict(&words))
	fmt.Println(findSubstring(s, words))
	fmt.Println(findSubstring(s2, words2))
}
