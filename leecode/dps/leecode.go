package dps

import (
	"fmt"

	"math"
	"strconv"
)

func lengthOfLongestSubstring(s string) int {
	f := func(sl []uint8, s uint8) (int, bool) {
		for i := 0; i < len(sl); i++ {
			if sl[i] == s {
				return i, true
			}
		}
		return -1, false
	}
	var max int = 0
	//us=append(us,s[0])
	//start:=1
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var subnum int
	for j := 0; j < len(s); j++ {
		var posrepeat int = 0
		us := make([]uint8, 0)
		us = append(us, s[j])

		if subnum > max {
			max = subnum
		}
		//fmt.Printf("sub is %v,max is %v\n",subnum,max)
		subnum = 0
		for start := j + 1; start < len(s); start++ {

			if x, ok := f(us, s[start]); ok {
				posrepeat = x
				break
			}
			//fmt.Printf("start is %v,us is %v\n",s[start],us)
			us = append(us, s[start])
			subnum++
		}
		j = j + posrepeat
	}
	return max + 1
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tag int=0
	var head =new(ListNode)
	var rt =new(ListNode)
	rt.Next=nil
	var rs1 =new(ListNode)
	var rs2	=new(ListNode)
	rs1=l1
	rs2=l2
	num1:=make([]int,0)
	num2:=make([]int,0)
	for ;rs1!=nil||rs2!=nil	 ;  {
		if rs1==nil {
			num1=append(num1,0)
		}else {
			num1=append(num1,rs1.Val)
			rs1=rs1.Next
		}
		if rs2 == nil {
			num2=append(num2,0)
		}else {
			num2=append(num2,rs2.Val)
			rs2=rs2.Next
		}
	}
	//fmt.Printf("num1 is %v\n,num2 is %v\n",num1,num2)
	head=rt
	for i:=0;i<len(num1) ;i++  {
		rthead:=new(ListNode)
		if (num1[i]+num2[i]+tag)>=10 {
			rt.Val=num1[i]+num2[i]-10+tag
			tag=1
		}else {
			rt.Val=num1[i]+num2[i]+tag
			tag=0
		}
		//fmt.Printf("in func val is %v,num1 is %v,num2 is %v \n",rt.Val,num1[i],num2[i])
		if i<len(num1)-1 {
			rt.Next=rthead
			rt=rthead
		}
		if i == len(num1)-1&&tag == 1 {
			rthead1:=new(ListNode)
			rthead1.Val=1
			rt.Next=rthead1
			rt=rthead1
		}
	}
	return head
}
*/

func addnums(nums []int, target int) []int {
	hm := make(map[int]int)
	res := make([]int, 2)
	for i, j := range nums {
		hm[j] = i
	}
	for i := 0; i < len(nums); i++ {
		if j, ok := hm[target-nums[i]]; ok && j != i {
			res[0] = nums[i]
			res[1] = nums[j]
			return res
		}
	}
	return nil
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	//inres:=make([]int,0)
	if len(nums) == 0 {
		return nil
	}

	findtwosum := func(nums []int, sum int) ([][]int, []int) {
		//var flag bool=false
		rnums := make([]int, 0)
		if len(nums) < 2 {
			fmt.Printf("nums is null\n")
			return nil, nil
		}
		res := make([][]int, 0)
		hasht := make(map[int]int)
		rmmutlihasht := make(map[int]bool)
		vrmmutlihasht := make(map[int]bool)
		for i, v := range nums {
			hasht[v] = i
			if v != (0 - sum) {
				rnums = append(rnums, v)
			}
		}

		for i, v := range nums {
			if vrmmutlihasht[v] {
				continue
			}
			//rmmutlihasht 将左边元素下标存入，保证只向右侧查找
			rmmutlihasht[i] = true
			if inode, ok := hasht[sum-v]; ok && !rmmutlihasht[inode] && !vrmmutlihasht[sum-v] {
				//rmmutlihasht[sum-v]=true
				if nums[inode] == v {
					rmmutlihasht[inode] = true
				}
				//fmt.Printf("res is %v\n",v)
				res = append(res, []int{(0 - sum), v, sum - v})
			}
			vrmmutlihasht[v] = true
		}
		return res, rnums
	}

	rnums := make([]int, len(nums))
	copy(rnums, nums)

	for len(rnums) > 2 {
		//rest:=findtwosum(append(nums[:inode],nums[inode+1:]...),(0-value))
		rest, r := findtwosum(rnums[1:], (0 - rnums[0]))
		if rest != nil {
			res = append(res, rest...)
		}

		rnums = r
	}
	return res
}

func threenums(nums []int) [][]int {
	//res1:=make([]int,3)

	res1 := make([]int, 3)
	rest := make([][]int, 0)
	f := func(nums []int, target int) []int {
		hm := make(map[int]int)
		res := make([]int, 2)
		for i, j := range nums {
			hm[j] = i
		}
		for i := 0; i < len(nums); i++ {
			if j, ok := hm[target-nums[i]]; ok && j != i {
				res[0] = nums[i]
				res[1] = nums[j]
				return res
			}
		}
		return nil
	}

	var i int
	//var tag int=0
	n := make([]int, len(nums))
	smap := make([]map[int]int, 0)
	for i = 0; i < len(nums)-1; i++ {
		var tag int = 0
		if cn := copy(n, nums); cn != len(n) {
			return nil
		}
		value := 0 - n[i]
		fslice := append(n[:i], n[i+1:]...)
		if res := f(fslice, value); res != nil {
			res1 = append(res, 0-value)
			for _, j := range smap {
				for _, key := range res1 {
					if _, ok := j[key]; ok {
						tag = 1
					}
				}
			}
			if tag == 0 {
				rest = append(rest, res1)
			}
			kmap := make(map[int]int)
			for kk, vv := range res1 {
				kmap[vv] = kk
			}
			smap = append(smap, kmap)

			//fmt.Printf("i is %v,value is %v,res is %v\n",i,value,res)

		}
		//fmt.Printf("num is %v,n is %v\n",nums,n)

	}
	return rest
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	yushu := (len(nums1) + len(nums2)) % 2
	var midden int
	var fmid float64
	var pos1 int = 0
	var pos2 int = 0
	if yushu != 0 {
		midpos := (len(nums1)+len(nums2))/2 + 1
		//fmt.Printf("mid pos is %v\n",midpos)
		for i := 0; i < midpos; i++ {
			if pos1 >= len(nums1) {
				midden = nums2[pos2]
				pos2++
			} else if pos2 >= len(nums2) {
				midden = nums1[pos1]
				pos1++
			} else {
				if nums1[pos1] <= nums2[pos2] {
					midden = nums1[pos1]
					pos1++
				} else {
					midden = nums2[pos2]
					pos2++
				}
			}
			fmid = float64(midden)
		}

	} else {
		var leftpos int = (len(nums1) + len(nums2)) / 2
		//var rightpos int=leftpos+1

		var (
			left  int
			right int
			pos1  int = 0
			pos2  int = 0
		)
		for i := 0; i < leftpos; i++ {
			if pos1 >= len(nums1) {
				left = nums2[pos2]
				pos2++
			} else if pos2 >= len(nums2) {
				left = nums1[pos1]
				pos1++
			} else {
				if nums1[pos1] <= nums2[pos2] {
					left = nums1[pos1]
					pos1++
				} else {
					left = nums2[pos2]
					pos2++
				}
			}
		}
		fmt.Printf("pos1 is %v,pos2 is %v,left is %v\n", pos1, pos2, left)
		if pos1 >= len(nums1) {
			right = nums2[pos2]
		} else if pos2 >= len(nums2) {
			right = nums1[pos1]
		} else if nums1[pos1] > nums2[pos2] {
			right = nums2[pos2]
		} else {
			right = nums1[pos1]
		}

		/*
					if nums1[pos1]>nums2[pos2]{
					right=nums2[pos2]
				}else {
					right=nums1[pos1]
				}


			pos1=0
			pos2=0
			for j := 0;j<rightpos;j++ {
				if pos1>=len(nums1){
					right=nums2[pos2]
					pos2++
				}else if pos2>=len(nums2) {
					right=nums1[pos1]
					pos1++
				}else {
					if nums1[pos1]<=nums2[pos2] {
						right=nums1[pos1]
						pos1++
					}else {
						right=nums2[pos2]
						pos2++
					}
				}
			}
		*/

		fmid = (float64(left) + float64(right)) / 2
		//fmt.Printf("l is %v,r is %v ,fmid is %v\n",left,right,fmid)
	}
	return fmid

}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var (
		max         int = 0
		nowl        int
		nowdouble   int
		plstrdouble string
		plstr       string
		maxplstr    string
	)
	for i := 0; i < len(s)-1; i++ {
		nowl = 0
		plstr = string(s[i])
		left := i - 1
		right := i + 1
		nowl = len(plstr)
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				plstr = string(s[left]) + plstr + string(s[right])
				left--
				right++
				nowl = len(plstr)
			} else {
				break
			}
		}
		if s[i] == s[i+1] {
			plstrdouble = string(s[i]) + string(s[i+1])
			left := i - 1
			right := i + 2
			nowdouble = 2
			for left >= 0 && right < len(s) {
				if s[left] == s[right] {
					plstrdouble = string(s[left]) + plstrdouble + string(s[right])
					left--
					right++
					nowdouble = len(plstrdouble)
				} else {
					break
				}
			}
			if nowdouble > nowl {
				plstr = plstrdouble
				nowl = nowdouble
			}
		}
		if nowl > max {
			max = nowl
			maxplstr = plstr
		}
		//fmt.Printf("max is %v,plstr is %v\n",maxplstr,plstr)
	}
	return maxplstr
}

func longestPalindrome1(s string) string {
	expandAroundCenter := func(s string, left, right int) int {
		ll := left
		rr := right
		for ll >= 0 && rr < len(s) && s[ll] == s[rr] {
			ll--
			rr++
		}
		return rr - ll - 1
	}
	var (
		start  int = 0
		end    int = 0
		len1   int
		len2   int
		slen   int = 1
		maxstr string
	)
	f := expandAroundCenter
	for i := 0; i < len(s); i++ {
		len1 = f(s, i, i)
		len2 = f(s, i, i+1)
		if len1 > len2 {
			slen = len1
		} else {
			slen = len2
		}
		if slen > end-start {
			start = i - (slen-1)/2
			end = i + slen/2
		}
		//fmt.Printf("start %v,end %v,slen is %v\n",start,end,slen)
	}

	maxstr = string(s[start])
	for j := start + 1; j <= end; j++ {

		maxstr = maxstr + string(s[j])
	}
	return maxstr

}

func maxArea(height []int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	f := min
	if len(height) == 1 {
		return 0
	}
	if len(height) == 2 {
		return f(height[0], height[1])
	}
	var (
		maxcontainer int = 0
		cont         int = 0
	)
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			cont = f(height[i], height[j]) * (j - i)
			if cont >= maxcontainer {
				maxcontainer = cont
			}
		}
	}
	return maxcontainer
}

func maxArea1(height []int) int {
	var l int = 0
	var r int = len(height) - 1
	var maxarea int
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for r > l {
		maxarea = max(maxarea, min(height[l], height[r])*(r-l))
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return maxarea
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefixstr := func(s1, s2 string) string {
		var prefix string
		for i := 0; i < len(s1) && i < len(s2); i++ {
			if s1[i] == s2[i] {
				prefix = prefix + string(s1[i])
			} else {
				break
			}
		}
		return prefix
	}
	var compre string
	compre = strs[0]
	for i := 1; i < len(strs); i++ {
		compre = prefixstr(compre, strs[i])
	}
	return compre
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	//fmt.Printf("max is %d\n",math.MaxInt32)
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	getreverse := func(xx int) int {
		var time int = 10
		var yushu int
		var v int
		var rever int = 0
		yushu = x % time
		rever = yushu
		v = x / time
		for v != 0 {
			fmt.Printf("yus is %d,v is %d,re is %d\n", yushu, v, rever)
			x = v
			yushu = x % time
			rever = rever*10 + yushu
			v = x / time
		}
		return rever
	}

	return getreverse(x)

}

func isPalindrome(x int) bool {

	if x == 0 {
		return true
	}
	getreverse := func(xx int) int {
		var time int = 10
		var yushu int
		var v int
		var rever int = 0
		yushu = x % time
		rever = yushu
		v = x / time
		for v != 0 {
			x = v
			yushu = x % time
			rever = rever*10 + yushu
			v = x / time
		}

		return rever
	}
	m := x
	//fmt.Printf("r is %d,x is %d\n",r,m)
	if m != getreverse(x) {
		return false
	}
	return true

}

func putstr(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("the %d value is %v\n", i, str[i])
	}
}

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	length := len(str)
	for i := 0; i < length; i++ {
		//fmt.Printf("i is %v\n",str[i])
		if str[i] == 32 {
			//fmt.Printf("nill\n")
			str = str[i+1:]
			i--
			length--
		}
		//putstr(str)
	}
	t, err := strconv.Atoi(str)
	if err != nil {
		//fmt.Printf("err is %v\n",err)
		return 0
	}
	return t

}
