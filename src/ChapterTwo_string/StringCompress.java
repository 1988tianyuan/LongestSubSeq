package ChapterTwo_string;

/**
 * @author Liu Geng liu.geng@navercorp.com
 * @date 2020/5/18 11:07
 */
public class StringCompress {
	
	public static String stringCompress(String input) {
		if (input == null || input.length() == 0) {
			return "";
		}
		int len = input.length();
		char[] chars = input.toCharArray();
		Character currentChar = null;
		int adder = 0;
		StringBuilder builder = new StringBuilder();
		for (int i = 0; i < len; i++) {
			if (currentChar == null) {
				currentChar = chars[i];
			}
			if (currentChar == chars[i]) {
				adder++;
			} else {
				builder.append(currentChar).append(adder);
				adder = 1;
				currentChar = chars[i];
			}
		}
		if (adder != 0) {
			builder.append(currentChar).append(adder);
		}
		return builder.toString();
	}
}
