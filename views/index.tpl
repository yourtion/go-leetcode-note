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
            <td>
                <a target="_blank" href="{{.Problem.Url}}">{{.Problem.Pid}}. {{.Problem.Name}}</a>
            </td>
            <td>{{date .Day "Y-m-d" }}</td>
            <td>
                <a target="_blank" href="{{.Submissions}}">查看</a>{{if $.login }} | <a href="/note/{{.Id}}">编辑</a>{{end}}
            </td>
        </tr>
        {{end}}
        </tbody>
    </table>
</div>