
class Solution {
    public int maxArea(int[] height) {
        int left=0;
        int right = height.length-1;
        //绝对值
        int max=Math.min(height[right],height[left])*(right-left);
        while (left<=right) {
            //如果右侧为短板，移动右侧
            if(height[left]>=height[right]){
                for (;left<=right;--right) {
                    int m = Math.min(height[right], height[left]) * (right - left);
                    if (max < m) {
                        max = m;
                        break;
                    }
                    if(height[left]<height[right]){
                        break;
                    }
                }
            }else{
                for (;left<=right;++left) {
                    int m = Math.min(height[right], height[left]) * (right - left);
                    if (max < m) {
                        max = m;
                        break;
                    }
                    if(height[left]>=height[right]){
                        break;
                    }
                }
            }
        }
        return max;
    }
}
