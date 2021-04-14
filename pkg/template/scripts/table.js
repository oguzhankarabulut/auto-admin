var vm = new Vue({
    el: '#app',
    data: data,
    components: 'side-menu',
    template: `
        <div class="page-wrapper">
            <side-menu :collections="['TEST', 'TEST2']"/>
            <div class="page-content">
                <h2 v-text="data.CollectionName.replaceAll('_',' ')"></h2>
                <div class="table-wrapper">
                    <table v-if="data.Documents">
                        <thead>
                            <tr>
                                <th v-for="k in dataKeys" :key="k">{{k}}</th>
                                <th></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="d in data.Documents" :key="d._id">
                                <td v-for="k in dataKeys" :key="k">{{d[k]}}</td>
                                <td><a :href="'/' + data.CollectionName + '?id=' + d._id">Detail</a></td>
                            </tr>
                        </tbody>
                    </table>
                    <div class="data-not-found" v-else>
                        <h1>🤔</h1>
                        <h3>Data Bulunamadı</h3>
                    </div>
                </div>
            </div>
        </div>
    `,
    computed:{
        dataKeys(){
            if(data.Documents){
                let keys =  Object.keys(data.Documents[0]);
                let returningKeys = keys.filter( item => item !== '__v');
                return returningKeys;
            }
        }
    },
})