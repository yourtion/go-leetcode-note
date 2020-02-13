<div id="index">
    <div class="pure-g">
        <div class="pure-u-4-5"><h2>总数：{{.count}}</h2></div>
        <div class="pure-u-1-5">
            <p>
            {{if .login}}
            <div class="pure-menu pure-menu-horizontal">
                <ul class="pure-menu-list">
                    <li class="pure-menu-item pure-menu-selected"><a href="/note" class="pure-menu-link">新建</a></li>
                    <li class="pure-menu-item pure-menu-has-children pure-menu-allow-hover">
                        <a href="#" id="menuLink1" class="pure-menu-link">导出</a>
                        <ul class="pure-menu-children">
                            <li class="pure-menu-item"><a href="/export" class="pure-menu-link">Blog</a></li>
                        </ul>
                    </li>
                </ul>
            </div>
            {{end}}
            </p>
        </div>
    </div>
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
                    <a target="_blank" href="{{.Problem.Url}}">{{.Problem.Pid}}
                        . {{.Problem.Name}}</a>{{if .Mark}}*{{end}}
                </td>
                <td>{{date .Day "Y-m-d" }}</td>
                <td>
                    <a target="_blank" href="{{.Submissions}}">查看</a>{{if $.login }} | <a
                        href="/note/{{.Id}}">编辑</a>{{end}}
                </td>
            </tr>
        {{end}}
        </tbody>
    </table>
</div>