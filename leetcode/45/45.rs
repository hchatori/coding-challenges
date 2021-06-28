struct Solution {}

impl Solution {
    pub fn jump(nums: Vec<i32>) -> i32 {
        if nums.len() == 1 {
            return 0;
        }

        // Create a vector that corresponds to nums, where the values in the
        // vector represent the farthest index that can be jumped to for a
        // given index in nums.
        let mut jump_position_index = vec![0; nums.len()];
        for i in 0..nums.len() {
            jump_position_index[i] = nums[i] + i as i32;
        }

        let mut jump_count = 0;
        let mut curr_max_jump_pos = 0;
        let mut next_max_jump_pos = 0;

        for i in 0..jump_position_index.len() {
            // Keep track of the largest index (farthest index) that can be
            // jumped to.
            if jump_position_index[i] > next_max_jump_pos {
                next_max_jump_pos = jump_position_index[i];
            }

            // Upon reaching the max index that can be jumped to given the
            // current jump, update the curr_max_jump_pos using the
            // next_max_jump_pos that has been kept track of above, and increment
            // the jump_count. If the current jump can reach the end
            // (curr_max_jump_pos >= goal), then return jump_count.
            if i as i32 == curr_max_jump_pos {
                curr_max_jump_pos = next_max_jump_pos;
                jump_count += 1;
                println!("{}", String::from("I'm here!"));
                if curr_max_jump_pos >= (nums.len() - 1) as i32 {
                    return jump_count;
                }
            }
        }

        return 0;
    }
}

fn main() {
    // 2
    // let nums = vec![2, 3, 1, 1, 4];

    // 2
    // let nums = vec![2, 3, 0, 1, 4];

    // 1
    // let nums = vec![1, 2];

    // 3
    // let nums = vec![5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0];

    // 4
    // let nums = vec![1, 1, 1, 1, 1];

    // 3
    let nums = vec![5, 4, 0, 1, 3, 6, 8, 0, 9, 4, 9, 1, 8, 7, 4, 8];

    println!("{}", Solution::jump(nums));
}
