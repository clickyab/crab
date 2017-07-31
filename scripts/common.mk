need_root :
	@[ "$(shell id -u)" -eq "0" ] || exit 1

not_root :
	@[ "$(shell id -u)" != "0" ] || exit 1