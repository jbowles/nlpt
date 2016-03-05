package tkz

/*
* Load some Constants for easier testing of data
 */

const (
	BadStr = "expect0.7rant7! Then I want to show Snow White and the Seven Dwarves. <=AndThe start of a new sentence. And\n then\n\nagain for One and NASA?"
	//44 words
	ThoreauOne = "I went to the & * % ^ $ # @ | -  woods because I wished to live deliberately, to front only the essential facts of life, and see if I could not learn what it had to teach, and not, when I came to die, discover that I had not lived."
	//30 words
	ThoreauTwo = "If one advances confidently in the direction of his dreams, and endeavors to live the life which he has imagined, he will meet with a success unexpected in common hours."
	//19 words
	ThoreauThree = "What lies behind us and what lies ahead of us are tiny matters compared to what lives within us."
)

var thoneByte = []byte(ThoreauOne)
var ththreeeByte = []byte(ThoreauThree)
var badstrByte = []byte(BadStr)
