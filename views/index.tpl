<div id="index">
    <h2>总数：{{.count}}</h2>
    <table class="pure-table pure-table-horizontal">
        <thead>
        <tr>
            <th>#</th>
            <th>题目</th>
            <th>时间</th>
            <th>操作</th>
        </tr>
        </thead>
        <tbody>
        {{range .data}}
        <tr>
            <td>{{.Id}}</td>
            <td>{{.Problem.Pid}}. {{.Problem.Name}}</td>
            <td>{{date .Day "Y-m-d" }}</td>
            {{if $.login }}
            <td><a href="/note/{{.Id}}">编辑</a></td>
            {{else}}
            <td>查看</td>
            {{end}}
        </tr>
        {{end}}
        </tbody>
    </table>
</div>