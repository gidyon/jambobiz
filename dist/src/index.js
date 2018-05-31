const app = new Vue({
    template: `
        <nav class="fw fh">
            <div class="fw fh">
                <ul class="fh fw f-f_rw">
                    <li class="inline-flex fh">
                        <div class="fh">Company</div>
                    </li>
                    <li class="inline-flex fh" v-for="linkInfo in links" :key="linkInfo.link" v-cloak>
                        <a 
                            :href="linkInfo.link" 
                            :class="[isActive(linkInfo.link) ? 'a-active':'', 'fh']"
                            @click="reload"
                            >{{linkInfo.name}}</a>
                    </li>
                </ul>
            </div>
        </nav>`,
    el: '#nav',
    computed: {
        isActive() {
            return (link) => window.location.href.includes(link, 0)
        }
    },
    methods: {
        reload: () => location.reload()
    },
    data() {
        return {
            links: [
                { name: "About", link: "/about" },
                { name: "Terms of Services", link: "/terms-of-service" },
                { name: "Privacy Policy", link: "/privacy-policy" }
            ]
        }
    }
})