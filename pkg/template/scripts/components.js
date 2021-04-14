new Vue();

Vue.component('side-menu', {
    props: ['collections'],
    template: `
        <div class="side-menu-wrapper">
             <a href="/dashboard">Auto Admin</a>
             <ul class="side-menu">
                <li v-for="c in collections">
                    <a :href="'/' + c">{{c}}</a>
                </li>
            </ul>
        </div>
    `
})