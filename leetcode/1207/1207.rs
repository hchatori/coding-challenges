use std::collections::HashMap;
use std::collections::HashSet;

fn main() {
	let arr = vec![1, 2, 2, 1, 1, 3];
	println!("{}", Solution::unique_occurrences(arr));
}

struct Solution {}

impl Solution {
    pub fn unique_occurrences(arr: Vec<i32>) -> bool {
		// Create a HashMap where they key is a number in arr and the value is
		// the frequency in which the number occurs in arr.
		let mut counts = HashMap::new();
		for num in arr {
			let val = counts.entry(num).or_insert(0);
			*val += 1;
		}

		// Create a HashSet and check the set before adding
		let mut freqs = HashSet::new();
		for (_, value) in &counts {
			if freqs.contains(value) {
				return false;
			}
			freqs.insert(value);
		}
		return true;
    }
}