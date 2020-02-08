<div id="note">
    <form class="pure-form pure-form-aligned" method="post" autocomplete="off">
        <label><input name="id" hidden="hidden" value="{{.note.Id}}"/></label>
        <fieldset class="pure-group">
            <legend>题目</legend>
            <input id="p-title" name="p-title" placeholder="题目" class="pure-u-1" type="text"
                   value="{{if .note }}{{.note.Problem.Pid}}. {{.note.Problem.Name}}{{end}}"
                   required="required"/>
            <input id="p-url" name="p-url" placeholder="题目URL" class="pure-u-1" type="url" value="{{.note.Problem.Url}}"
                   required="required"/>
        </fieldset>

        <fieldset class="pure-group">
            <legend>题解</legend>
            <input id="submissions" name="submissions" placeholder="答案URL" class="pure-u-1" type="url"
                   value="{{.note.Submissions}}" required="required"/>
            <textarea id="solution" name="solution" placeholder="思路" cols="10" class="pure-u-1" required="required">{{.note.Solution}}</textarea>

            <textarea id="rethink" name="rethink" placeholder="反思" cols="3" class="pure-u-1" required="required">{{.note.Rethink}}</textarea>

            <textarea id="harvest" name="harvest" placeholder="收获" cols="2"
                      class="pure-u-1">{{.note.Harvest}}</textarea>

        </fieldset>
        <div class="pure-g">
            <div class="pure-u-4-5 pure-u-md">
                <label for="mark" class="pure-checkbox">
                    <input id="mark" name="mark" type="checkbox" value="{{.note.Mark}}"> 标记 </input>
                </label>
            </div>
            <div class="pure-u-1-5 pure-u-md">
                <button id="n-submit" type="submit" value="submit" class="pure-button pure-button-primary">提交</button>
            </div>
        </div>
    </form>
</div>