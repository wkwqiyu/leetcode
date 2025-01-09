class Solution {
    public void moveZeroes(int[] nums) {
		int[] result = new int[nums.length];
		int j = 0;
		for(int i=0;i<nums.length;i++){
			if(nums[i]!=0){
				if(j==i){
					j++;
				}else{
					nums[j++]=nums[i];
				}
			}
		}
		for(;j<nums.length;j++){
			nums[j]=0;
		}
    }
}