<script type="text/ecmascript-6">
    import StylesMixin from './../../mixins/entriesStyles';
    export default {
        mixins: [
            StylesMixin,
        ],
    }
</script>

<template>
<div>
        <index-screen title="Processes" resource="worker_pools">
        <tr slot="table-header">
            <th scope="col">Host</th>
            <th scope="col">Concurrency</th>
            <th scope="col">Pid</th>
            <th scope="col">Workers</th>
             <th scope="col">Started</th>
             <th scope="col">Checkin</th>
            <th scope="col"></th>
        </tr>


        <template slot="row" slot-scope="slotProps">
            <td :title="slotProps.entry.host" class="table-fit pr-0">
                {{truncate(slotProps.entry.host, 60)}}
                <div><small class="text-muted text-xs">{{slotProps.entry.worker_pool_id}}</small></div>    
            </td>
            <td class="table-fit">
                <span class="badge font-weight-light" :class="'badge-dark'">
                    {{slotProps.entry.concurrency}}
                </span>
            </td>

        
            <td class="table-fit">
                <span class="badge font-weight-light">
                    {{slotProps.entry.pid}}
                </span>
            </td>

            <td class="table-fit">
                <span class="badge font-weight-light" :class="'badge-info'">
                    {{slotProps.entry.worker_ids.length }}
                </span>
            </td>

            <td class="table-fit" :data-timeago="slotProps.entry.started_at" :title="slotProps.entry.started_at * 1000">
                {{timeAgo(slotProps.entry.started_at * 1000)}}
            </td>

             <td class="table-fit" :data-timeago="slotProps.entry.heartbeat_at * 1000" :title="slotProps.entry.heartbeat_at * 1000">
                {{timeAgo(slotProps.entry.heartbeat_at * 1000)}}
            </td>

            <td class="table-fit">
                <router-link :to="{name:'processes-preview', params:{id: slotProps.entry.pid}}" class="control-action">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 22 16">
                        <path d="M16.56 13.66a8 8 0 0 1-11.32 0L.3 8.7a1 1 0 0 1 0-1.42l4.95-4.95a8 8 0 0 1 11.32 0l4.95 4.95a1 1 0 0 1 0 1.42l-4.95 4.95-.01.01zm-9.9-1.42a6 6 0 0 0 8.48 0L19.38 8l-4.24-4.24a6 6 0 0 0-8.48 0L2.4 8l4.25 4.24h.01zM10.9 12a4 4 0 1 1 0-8 4 4 0 0 1 0 8zm0-2a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"></path>
                    </svg>
                </router-link>
            </td>
        </template>
    </index-screen>

    <div class="mt-4">
        
    </div>
        <index-screen title="Busy processes" resource="busy_workers">
        <tr slot="table-header">
            <th scope="col">Host</th>
            <th scope="col">Concurrency</th>
            <th scope="col">Pid</th>
            <th scope="col">Workers</th>
             <th scope="col">Started</th>
             <th scope="col">Checkin</th>
            <th scope="col"></th>
        </tr>


        <template slot="row" slot-scope="slotProps">
            <td :title="slotProps.entry.host" class="table-fit pr-0">
                {{truncate(slotProps.entry.host, 60)}}
                <div><small class="text-muted text-xs">{{slotProps.entry.worker_pool_id}}</small></div>    
            </td>
            <td class="table-fit">
                <span class="badge font-weight-light" :class="'badge-dark'">
                    {{slotProps.entry.concurrency}}
                </span>
            </td>

        
            <td class="table-fit">
                <span class="badge font-weight-light">
                    {{slotProps.entry.pid}}
                </span>
            </td>

            <td class="table-fit">
                <span class="badge font-weight-light" :class="'badge-info'">
                    {{slotProps.entry.worker_ids.length }}
                </span>
            </td>

            <td class="table-fit" :data-timeago="slotProps.entry.started_at" :title="slotProps.entry.started_at * 1000">
                {{timeAgo(slotProps.entry.started_at * 1000)}}
            </td>

             <td class="table-fit" :data-timeago="slotProps.entry.heartbeat_at * 1000" :title="slotProps.entry.heartbeat_at * 1000">
                {{timeAgo(slotProps.entry.heartbeat_at * 1000)}}
            </td>

            <td class="table-fit">
                <router-link :to="{name:'processes-preview', params:{id: slotProps.entry.pid}}" class="control-action">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 22 16">
                        <path d="M16.56 13.66a8 8 0 0 1-11.32 0L.3 8.7a1 1 0 0 1 0-1.42l4.95-4.95a8 8 0 0 1 11.32 0l4.95 4.95a1 1 0 0 1 0 1.42l-4.95 4.95-.01.01zm-9.9-1.42a6 6 0 0 0 8.48 0L19.38 8l-4.24-4.24a6 6 0 0 0-8.48 0L2.4 8l4.25 4.24h.01zM10.9 12a4 4 0 1 1 0-8 4 4 0 0 1 0 8zm0-2a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"></path>
                    </svg>
                </router-link>
            </td>
        </template>
    </index-screen>
</div>
</template>