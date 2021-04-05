use std::collections::HashMap;
use std::fs::File;
use std::io::{ BufReader, BufRead, Error};

fn is_sum_possible(sum_number: i64, numbers: &[i64]) -> bool {
	let mut numbers_dict:HashMap<i64, i64>  = HashMap::new(); /* could be replaced by an hashset */

	for num in numbers.iter() {
		if numbers_dict.contains_key(&num) {
			return true;
		} 
		else {
			let diff = sum_number-num;
                        numbers_dict.insert(diff, *num);
		}
	}

	false
}

fn solve(numbers: &[i64]) -> Result<i64, &'static str> {
	let num_len = numbers.len();

	for index in 25..num_len {
		let prev_numbers = &numbers[index-25..index];
		let current_number = numbers[index];

		if !is_sum_possible(current_number, prev_numbers) {
			return Ok(current_number);
		}

	}

	Err("Number not found")
}

fn main() -> Result<(), Error> {
    let path = "input";

    let input = File::open(path)?;
    let buffered = BufReader::new(input);

    let mut numbers = Vec::new();

    for line in buffered.lines() {
	let num = line?.parse::<i64>().unwrap();
    	numbers.push(num);
    }

    let solution = solve(&numbers).unwrap();
    println!("Solution is: {:?}",solution);

    Ok(())
}
