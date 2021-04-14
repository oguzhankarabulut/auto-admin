var vm = new Vue({
    el: '#app',
    data: data,
    template: `
    <div class="page-content">
        <h2>{{data.CollectionName.replaceAll('_',' ')}} Detail</h2>
        <div class="table-wrapper">
            <table>
                <thead>
                    <tr>
                        <th>Key</th>
                        <th>Value</th>
                    </tr>
                </thead>
                <tbody>
                    <tr  v-for="k in dataKeys" :key="k">
                        <td>
                            {{k}}
                        </td>
                          <td>
                            {{data.Detail[k]}}
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    `,
    computed:{
        dataKeys(){
            let keys =  Object.keys(data.Detail);
            let returningKeys = keys.filter( item => item !== '__v');
            return returningKeys;
        }
    },
});