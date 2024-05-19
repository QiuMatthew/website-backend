import sys

# Check if enough arguments are provided
if len(sys.argv) != 4:
    print("Error: Missing arguments. Expected 3 arguments.")
    sys.exit(1)

# Retrieve arguments
g = sys.argv[1]
h = sys.argv[2]
p = sys.argv[3]

# Use the arguments (for demonstration, we just print them)
print(f"g: {g}, h: {h}, p: {p}")
