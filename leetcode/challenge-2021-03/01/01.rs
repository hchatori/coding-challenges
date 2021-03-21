use std::collections::HashSet;

impl Solution {

    pub fn distribute_candies(candy_type: Vec<i32>) -> i32 {
        let max_allowed = candy_type.len() / 2;
        let mut c_types = HashSet::new();

        for t in candy_type.iter() {
            c_types.insert(*t);
        }
        
        if c_types.len() < max_allowed {
            return c_types.len() as i32;
        }

        return max_allowed as i32;
    }
}

struct Solution {}

fn main() {
    let candy_type = vec![0, 0, 0, 4];
    println!("{}", Solution::distribute_candies(candy_type));
}