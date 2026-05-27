do_syscall:
	push rbp
	mov rbp, rsp

    mov rax, rdi
    mov rdi, rsi
    mov rsi, rdx
    mov rdx, r10
    syscall

	mov rsp, rbp
	pop rbp
    ret
