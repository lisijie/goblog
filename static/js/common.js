var hscms = {
	'MSG_NONE': 0,
	'MSG_OK': 1,
	'MSG_ERR': 2,
	'MSG_LOGIN': 3,
	'ajax': {
		'getJSON':function(url, cb) {
			$.getJSON(url, function (out) {
				if (out.ret == 1) {
					cb(out.data);
				} else if (out.ret == 2) {
					alert(out.msg);
				}
			});
		}
	},
	'comment': {
		'support':function(id, o) {
			hscms.ajax.getJSON('?mod=comment&act=digg&commentid='+id, function(data) {
				$(o).before('<a>\u5df2\u9876<span class="num">['+data.up+']</span></a>').remove();
			});
		}
	}
};
