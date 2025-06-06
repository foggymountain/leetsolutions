package main;

public class Dupes {

	public static void main(String[] args) {
		int[] a = { 0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 4 };
		removeDuplicates(a);
	}

	public static int removeDuplicates(int[] nums) {
		if (nums.length == 0)
			return 0;

		int i = 0; 
		for (int j = 1; j < nums.length; j++) {
			if (nums[j] != nums[i]) {
				i++;
				nums[i] = nums[j]; 
			}
		}
		return i + 1; 
	}
}
