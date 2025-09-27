define i32 @main() {
entry:
	%0 = alloca i32
	store i32 0, i32* %0
	%1 = alloca i32
	store i32 1, i32* %1
	%2 = alloca i32
	store i32 2, i32* %2
	%3 = alloca i32
	store i32 10, i32* %3
	%4 = alloca i32
	store i32 0, i32* %4
	br label %while.cond-1

while.cond-1:
	%5 = load i32, i32* %2
	%6 = load i32, i32* %3
	%7 = icmp slt i32 %5, %6
	br i1 %7, label %while.body-1, label %main-1

while.body-1:
	%8 = load i32, i32* %1
	store i32 %8, i32* %4
	%9 = load i32, i32* %0
	%10 = load i32, i32* %1
	%11 = add i32 %9, %10
	store i32 %11, i32* %1
	%12 = load i32, i32* %4
	store i32 %12, i32* %0
	%13 = load i32, i32* %2
	%14 = add i32 %13, 1
	store i32 %14, i32* %2
	br label %while.cond-1

main-1:
	%15 = load i32, i32* %1
	ret i32 %15
}
