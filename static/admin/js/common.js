
function checkall(name, obj) {
	$(":checkbox[name='"+name+"']").each(function(o) {
		$(this).prop('checked', obj.checked);
	});
}

function submitAct(mod, act) {
	var action = "?mod="+mod+"&act="+act+"&t="+Date.now();
	if (arguments.length == 3) {
		action += "&"+arguments[2];
	}
	if (confirm('确定要执行该操作吗？')) {
		$('#mainform').attr('action', action).submit();
	}
}

function del_confirm() {
	return confirm('一旦删除将不可恢复，确定吗？');
}