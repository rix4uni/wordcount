import json

def process_subdomains(parts):
    main_domain = parts[-2] + '.' + parts[-1]  # e.g., 'telenet.be' or 'dell.com'
    unique_parts = []
    for part in parts[:-2]:  # exclude the last two parts (main domain)
        if part not in main_domain:
            unique_parts.append(part)
    return main_domain, unique_parts

# Open the file in read mode
with open('subs.txt', 'r') as file:
    # Read all lines from the file
    lines = file.readlines()

# Dictionary to store word counts
word_counts = {}

# Dictionary to track unique parts for each main domain
domains_parts = {}

# Loop through each line
for line in lines:
    # Strip leading/trailing whitespace from the line
    parts = line.strip().split('.')
    # Check if the line contains more than one dot
    if len(parts) > 2:
        # Process the subdomains
        main_domain, processed_parts = process_subdomains(parts)
        if main_domain not in domains_parts:
            domains_parts[main_domain] = set()
        # Add words to the dictionary with their counts
        for part in processed_parts:
            if part not in domains_parts[main_domain]:
                domains_parts[main_domain].add(part)
                if part not in word_counts:
                    word_counts[part] = 0
                word_counts[part] += 1

# Sort the word counts dictionary by counts in descending order
sorted_word_counts = dict(sorted(word_counts.items(), key=lambda item: item[1], reverse=True))

# Save the sorted dictionary as a JSON file
with open('best-dns-wordlist.json', 'w') as json_file:
    json.dump(sorted_word_counts, json_file, indent=4)

print("best-dns-wordlist JSON file created successfully.")

# zcat all-subdomains.txt.gz | grep -a '\..*\.' | sed 's/\.[^.]*\.[^.]*$//' | unew -q subs.txt