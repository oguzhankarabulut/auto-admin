var vm = new Vue({
    el: '#app',
    data: data,
    template: `
    <div class="page-content">
        <h2 v-text="data.Title.replaceAll('_', ' ')"></h2>
        <ul class="side-menu">
            <li v-for="c in data.Collections">
                <a :href="'/' + c" v-text="c.replaceAll('_', ' ')"></a>
            </li>
        </ul>
    </div>
    `,
    computed:{

    },
});
console.log(vm);