with open("input") as fn:
	matching_pwds = 0
	for line in fn.readlines():
		parts = line.split(":")
		password = parts[1].strip()
		policy = parts[0]
		policy_parts = policy.split(" ")
		char_to_look = policy_parts[1]
		num_range = policy_parts[0].split("-")
		min_occ = int(num_range[0])
		max_occ = int(num_range[1])
		cnt = 0
		for p in password:
			if p == char_to_look:
				cnt += 1
		if cnt >= min_occ and cnt <= max_occ:
			matching_pwds += 1

print(matching_pwds)
		
			
