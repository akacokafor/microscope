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
        <index-screen title="Retry Jobs" resource="retry_jobs">
        <tr slot="table-header">
            <th scope="col">Info</th>
            <th scope="col">Fails</th>
            <th scope="col">Failed</th>
            <th scope="col">Retry At</th>
            <th scope="col"></th>
        </tr>


        <template slot="row" slot-scope="slotProps">
            <td :title="slotProps.entry.job_name" class="table-fit pr-0">
                {{truncate(slotProps.entry.name, 60)}}
                <div>
                    <small class="text-danger">{{ slotProps.entry.err }}</small>
                </div>
            </td>
            <td class="table-fit">
                <span class="badge font-weight-light" :class="'badge-dark'">
                    {{slotProps.entry.fails}}
                </span>
            </td>

        
                   <td class="table-fit" :data-timeago="slotProps.entry.failed_at * 1000" :title="slotProps.entry.failed_at * 1000">
                {{timeAgo(slotProps.entry.failed_at * 1000)}}
            </td>

                 <td class="table-fit" :data-timeago="slotProps.entry.retry_at * 1000" :title="slotProps.entry.retry_at * 1000">
                {{timeAgo(slotProps.entry.retry_at * 1000)}}
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