<!DOCTYPE html>
	<html>
	<meta http-equiv="refresh" content="5" >
	<style>
	table{
		border:1px solid gray;
		width:45%;
		
	}
	.tbldata th{
		border:1px solid black;
		background: linear-gradient(to bottom,  #1e5799 0%,#2989d8 50%,#207cca 51%,#7db9e8 100%);
		padding:10px;

	}
	.tbldata td{
		border:1px solid black;
		padding:10px;
	}
	</style>
		<head>
			<meta charset="UTF-8">
			<title>Data Angin & Air</title>
		</head>
		<body>
			<table class="tbldata" cellspacing="0">
			<tr>
			<th colspan="3"><h2 style="color:white">WATER</h2></th>
			</tr>
			<tr>
				<th rowspan="2">
				<img src="https://media.tenor.com/images/59da2ec464f49cce2bab7105586cd844/tenor.gif" width="200px">
				</th>

				<th>
				Value
				</th>
				<th>
				Status
				</th>
			</tr>
			<tr>
			{{$wtr:=.Water}}
				<td>{{.Water}} m</td>
				{{if lt $wtr 30}}
				<td><span style="color:green">Aman</span></td>
				{{else if and (ge $wtr 30 ) (le $wtr 60)}}
				<td><span style="color:orange">Siaga</span></td>
				{{else if gt $wtr 60}}
				<td><span style="color:red">Bahaya</span></td>
				{{end}}
			</tr>
			</table>
			<table class="tbldata" cellspacing="0">
			<tr>
			<th colspan="3"><h2 style="color:white">WIND PRESSURE</h2></th>
			</tr>
			<tr>
				<th rowspan="2">
				<img src="https://thumbs.gfycat.com/WhoppingInnocentBaboon-small.gif" width="200px">
				</th>

				<th>
				Value
				</th>
				<th>
				Status
				</th>
			</tr>
			<tr>{{$wnd:=.Wind}}
				<td>{{.Wind}} m/s</td>
				{{if lt $wnd 25}}
				<td><span style="color:green">Aman</span></td>
				{{else if and (ge $wnd 25 ) (le $wnd 50)}}
				<td><span style="color:orange">Siaga</span></td>
				{{else if gt $wnd 50}}
				<td><span style="color:red">Bahaya</span></td>
				{{end}}
			</tr>

			</table>
		</body>
	</html>