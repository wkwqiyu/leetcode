class Solution {
    public int longestConsecutive(int[] nums) {

		if(nums.length==0){
			return 0;
		}
		Set<Integer> m = new HashSet<>();
		//给nums排序
		for(int i =0;i<nums.length;i++){
			m.add(nums[i]);
		}
		int maxl = 0;
		for(Integer i :m) {
		//降低运算复杂度的关键
			if(m.contains(i-1)){
				continue;
			}
			int l = 0;
			while(m.contains(i++)){
				l++;
			}
			maxl = Math.max(maxl,l);
		}
		return maxl;
    }
}