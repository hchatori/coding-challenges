struct Solution {}

impl Solution {
    pub fn predict_party_victory(senate: String) -> String {
        let mut senators: Vec<char> = senate.chars().collect();
        let mut victory = false;
        let mut victor = String::from("");

        if senators.len() == 1 {
            if senators[0] == 'R' {
                victor = String::from("Radiant");
                return victor;
            } else {
                victor = String::from("Dire");
                return victor;
            }
        }

        while !victory {
            let mut i = 0;
            let mut j = i + 1;
            let mut second_pass = false;

            while i < senators.len() {
                // If 'X', then skip
                if senators[i] == 'X' {
                    i += 1;
                    continue;
                }

                let mut reset = false;
                while j < senators.len() {
                    // If 'R', then find next 'D' and change 'D' to 'X'
                    // or if there are none to change, declare victory.
                    if senators[i] == 'R' {
                        if senators[j] == 'D' {
                            senators[j] = 'X';
                            break;
                        } else if j == senators.len() - 1 {
                            if !second_pass {
                                reset = true;
                            } else {
                                victor = String::from("Radiant");
                                victory = true;
                            }
                        }
                    }
                    // If 'D', then find next 'R' and change 'R' to 'X'
                    // or if there are none to change, declare victory.
                    if senators[i] == 'D' {
                        if senators[j] == 'R' {
                            senators[j] = 'X';
                            break;
                        } else if j == senators.len() - 1 {
                            if !second_pass {
                                reset = true;
                            } else {
                                victor = String::from("Dire");
                                victory = true;
                            }
                        }
                    }
                    j += 1;
                }

                if reset {
                    j = 0;
                    second_pass = true;
                    continue;
                }
                i += 1;
            }
        }

        return victor;
    }
}

fn main() {
    // Radiant
    // let senate = String::from("RD");

    // Dire
    // let senate = String::from("RDD");

    // Dire
    // let senate = String::from("DDRRR");

    // Dire
    let senate = String::from("DRRDRDRDRDDRDRDRD");

    // Radiant
    // let senate = String::from("DRRR");

    println!("{}", Solution::predict_party_victory(senate));
}
