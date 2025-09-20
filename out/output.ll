define i32 @main() {
entry:
	%0 = alloca i32
	store i32 5, i32* %0
	br label %while.cond-1

while.cond-1:
	%1 = load i32, i32* %0
	%2 = icmp slt i32 %1, 10
	br i1 %2, label %while.body-1, label %main-1

while.body-1:
	%3 = load i32, i32* %0
	%4 = add i32 %3, 1
	store i32 %4, i32* %0
	br label %while.cond-1

main-1:
	%5 = load i32, i32* %0
	ret i32 %5
}
