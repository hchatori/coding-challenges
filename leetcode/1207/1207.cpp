#include <iostream>
#include <map>
#include <unordered_set>
#include <vector>

class Solution {
public:
    bool uniqueOccurrences(std::vector<int>& arr) {
		// Create a map where they key is a number in arr and the value is
		// the frequency in which the number occurs in arr.
		std::map<int, int> counts;
        for (const auto& num : arr) {
			if (counts.find(num) == counts.end()) {
				counts[num] = 1;
			} else {
				counts[num] += 1;
			}
		}

		// Create an unordered set and check the unordered set before adding
		std::unordered_set<int> freqs;
		for (auto const& p : counts) {
			if (freqs.find(p.second) != freqs.end()) {
				return false;
			}
			freqs.insert(p.second);
		}
		return true;
    }
};

int main() {
	Solution s;
	std::vector<int> arr {1, 2, 2, 1, 1, 3};
	std::cout << s.uniqueOccurrences(arr) << std::endl;
	return 0;
}