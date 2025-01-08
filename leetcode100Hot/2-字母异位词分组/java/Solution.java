class Solution {
    public static List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> result = new ArrayList<>();
        Map<Integer,List<String>> m = new HashMap<>();
        for(String str:strs){
            char[] cs = str.toCharArray();
            Arrays.sort(cs);
            //排序后算出特征
            int i = Arrays.hashCode(cs);
            //根据哈希表存储相同特征的字符串
            if(m.containsKey(i)){
                m.get(i).add(str);
            }else{
                List<String> stringList = new ArrayList<>();
                stringList.add(str);
                m.put(i,stringList);
            }
        }
        for(Map.Entry<Integer,List<String>> entry:m.entrySet()){
            result.add(entry.getValue());
        }
        return result;
    }
}