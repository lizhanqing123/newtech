update draw_device a
set LayoutId = (                                                                                                                         
	select LayoutId
	from draw_layout
	where draw_layout.iBtsId = a.iBtsId and draw_layout.Floor = a.Position)
where LayoutId = 300008

update project inner join (select address_code from mixtubake.project ) as b ON project.ProjectName = b.ProjectName
SET project.address_code = b.address_code
