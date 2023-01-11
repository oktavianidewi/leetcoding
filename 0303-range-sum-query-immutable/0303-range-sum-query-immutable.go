type NumArray struct {
    Nums []int
}


func Constructor(nums []int) NumArray {
    x := NumArray{
        Nums:nums,
    }
    return x
}


func (this *NumArray) SumRange(left int, right int) int {
    sum:=0
    for i := left ; i<=right;i++{
        sum +=this.Nums[i]
    }
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */