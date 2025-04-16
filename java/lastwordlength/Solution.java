public class Solution {

	public static void main(String[] args) {
		LastWord lw = new LastWord();
		lw.lengthOfLastWord("Hello World");
	}

    public int lengthOfLastWord(String s) {
		
        String p = s.trim();
    
        int end = p.length();
        if (end == 1) return 1;

		int start = end;
		while (start > 0 && !" ".equals(p.substring(start-1, start))) {
			start--;
		}
		String k = p.substring(start, end);
		return k.length();
	
    }
}
