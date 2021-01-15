package main

import "testing"

func TestDecompress(t *testing.T) {
	var compresseds = [][]byte{
		{36, 0, 2, 1, 6, 2, 35, 3, 2, 0, 38, 33, 34, 32, 39, 0, 1, 6, 2, 35, 1, 1, 2, 6, 195, 0, 18, 8, 32, 0, 0, 6, 1, 0, 3, 3, 2, 133, 0, 27, 2, 2, 3, 35, 195, 0, 15, 1, 0, 6, 65, 0, 1, 133, 0, 46, 98, 1, 3, 3, 33, 32, 38, 36, 0, 196, 0, 74, 132, 0, 46, 1, 2, 3, 133, 0, 35, 134, 0, 82, 134, 0, 68, 133, 0, 15, 4, 3, 6, 4, 2, 1, 39, 2, 2, 6, 38, 35, 133, 0, 79, 0, 4, 41, 2, 0, 6, 133, 0, 137, 131, 0, 2, 3, 1, 2, 2, 5, 202, 0, 169, 14, 5, 5, 0, 5, 3, 5, 2, 2, 0, 5, 0, 2, 5, 0, 37, 196, 0, 183, 45, 5, 1, 37, 37, 143, 0, 197, 133, 0, 200, 3, 0, 0, 5, 2, 204, 0, 234, 195, 0, 242, 131, 0, 183, 131, 1, 5, 132, 0, 236, 0, 0, 197, 0, 201, 52, 0, 38, 3, 0, 2, 132, 1, 50, 2, 3, 3, 0, 202, 1, 46, 35, 2, 136, 0, 17, 138, 1, 65, 39, 0, 255},
		{0, 0, 107, 0, 198, 0, 6, 0, 0, 109, 12, 197, 0, 24, 109, 26, 197, 0, 44, 110, 40, 195, 0, 63, 111, 55, 196, 0, 83, 109, 71, 196, 0, 104, 1, 0, 0, 109, 85, 196, 0, 123, 1, 0, 0, 108, 99, 197, 0, 143, 34, 0, 105, 112, 37, 0, 2, 122, 123, 0, 105, 124, 6, 0, 124, 0, 134, 135, 0, 0, 109, 136, 5, 139, 138, 150, 151, 0, 0, 113, 152, 1, 0, 0, 113, 170, 1, 0, 0, 110, 188, 4, 0, 203, 204, 0, 0, 110, 205, 4, 0, 220, 221, 0, 0, 110, 222, 2, 0, 223, 237, 39, 0, 100, 238, 200, 1, 71, 38, 0, 99, 243, 36, 0, 98, 247, 255},
	}
	var origins = [][]byte{
		{0, 0, 0, 0, 0, 1, 6, 2, 3, 3, 3, 3, 0, 38, 33, 32, 32, 32, 0, 0, 0, 0, 0, 0, 0, 0, 6, 2, 1, 1, 1, 1, 2, 6, 0, 32, 32, 32, 32, 0, 0, 6, 1, 0, 3, 3, 2, 2, 1, 1, 1, 1, 2, 2, 3, 35, 32, 33, 38, 0, 0, 6, 0, 1, 0, 1, 2, 2, 1, 1, 1, 1, 1, 2, 3, 3, 33, 32, 38, 0, 0, 0, 0, 0, 3, 2, 1, 1, 1, 2, 2, 1, 1, 1, 2, 3, 32, 32, 32, 32, 0, 0, 0, 0, 3, 2, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 32, 32, 32, 0, 0, 0, 3, 6, 4, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 6, 38, 35, 0, 0, 0, 0, 0, 3, 4, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 6, 35, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 5, 5, 2, 2, 1, 1, 0, 0, 0, 0, 0, 0, 5, 5, 0, 5, 3, 5, 2, 2, 0, 5, 0, 2, 5, 0, 37, 0, 5, 5, 0, 0, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 37, 37, 5, 5, 0, 0, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 0, 5, 2, 5, 0, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 2, 5, 0, 0, 0, 5, 3, 5, 5, 3, 5, 5, 5, 5, 5, 0, 0, 0, 5, 0, 0, 5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 2, 3, 3, 2, 3, 3, 3, 3, 0, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 5, 4, 3, 2, 1, 0, 0, 0, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 15, 14, 13, 12, 0, 0, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 29, 28, 27, 26, 0, 0, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 42, 41, 40, 0, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 58, 57, 56, 55, 0, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 74, 73, 72, 71, 0, 0, 0, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 86, 85, 0, 0, 0, 0, 0, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 99, 0, 0, 0, 0, 0, 0, 0, 0, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 0, 0, 0, 0, 0, 0, 122, 123, 0, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 0, 124, 0, 134, 135, 0, 0, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 139, 138, 150, 151, 0, 0, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 0, 0, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 0, 0, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 0, 203, 204, 0, 0, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 0, 220, 221, 0, 0, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 0, 223, 237, 0, 0, 0, 0, 0, 0, 0, 0, 238, 239, 240, 241, 242, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 243, 244, 245, 246, 0, 0, 0, 0, 0, 247, 248, 249},
	}

	for i := 0; i < len(compresseds); i++ {
		compressed := compresseds[i]
		actual := Decompress(compressed)

		original := origins[i]
		if len(original) != len(actual) {
			t.Fatalf("TestRun failed: wrong size of decompressed data. want=%d, got=%d", len(original), len(actual))
		}

		for off, b := range original {
			if b != actual[off] {
				t.Fatalf("TestRun failed: offset:%d's data is wrong in decompressed data. want=0x%02x, got=%02x", off, b, actual[off])
			}
		}
	}
}