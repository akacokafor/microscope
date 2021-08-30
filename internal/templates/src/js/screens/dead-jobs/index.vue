<script type="text/ecmascript-6">
    import StylesMixin from './../../mixins/entriesStyles';
    export default {
        mixins: [
            StylesMixin,
        ],
        data() {
            return {
                modalContent: null,
            }
        },
        methods: {
            showQueueDetails(entry) {
            this.modalContent = entry;
                $('#queueJobInformation').modal({show: true})
            },
        }
    }
</script>

<template>
<div>
        <index-screen title="Dead Jobs" resource="dead_jobs">
        <tr slot="table-header">
            <th scope="col">Name</th>
            <th scope="col">Fails</th>
            <th scope="col">Failed</th>
            <th scope="col">Died</th>
            <th scope="col"></th>
        </tr>


        <template slot="row" slot-scope="slotProps">
          <td :title="slotProps.entry.name" class="table-fit pr-0">
                {{truncate(slotProps.entry.name, 60)}}
                <div :title="slotProps.entry.err">
                    <small class="text-danger">{{ truncate(slotProps.entry.err, 100) }}</small>
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

            <td class="table-fit" :data-timeago="slotProps.entry.died_at * 1000" :title="slotProps.entry.died_at * 1000">
                {{timeAgo(slotProps.entry.died_at * 1000)}}
            </td>


            <td class="table-fit">
                <a href="#" @click.prevent="showQueueDetails(slotProps.entry)" class="control-action">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 22 16">
                        <path d="M16.56 13.66a8 8 0 0 1-11.32 0L.3 8.7a1 1 0 0 1 0-1.42l4.95-4.95a8 8 0 0 1 11.32 0l4.95 4.95a1 1 0 0 1 0 1.42l-4.95 4.95-.01.01zm-9.9-1.42a6 6 0 0 0 8.48 0L19.38 8l-4.24-4.24a6 6 0 0 0-8.48 0L2.4 8l4.25 4.24h.01zM10.9 12a4 4 0 1 1 0-8 4 4 0 0 1 0 8zm0-2a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"></path>
                    </svg>
                </a>
            </td>
        </template>
    </index-screen>

    <div class="modal fade" id="queueJobInformation" tabindex="-1" aria-labelledby="queueJobInformation" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
        <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">Queue Details</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
            </button>
        </div>
        <div class="modal-body">
            <pre>{{ modalContent }}</pre>
        </div>
        <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
        </div>
        </div>
    </div>
</div>

</div>
</template>