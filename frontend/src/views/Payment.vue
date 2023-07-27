<script>
import axios from 'axios';

export default {
  data() {
    return {
      pay: null,
    };
  },
  created() {
    const orderId = this.$route.params.orderId; 
    axios.get(`http://localhost:8080/pay/${orderId}`, { responseType: 'arraybuffer' })
      .then(response => {
            const blob = new Blob([response.data],{type:'image/png'})
            this.pay = URL.createObjectURL(blob)
      })
      .catch(error => {
        console.error(error);
      });
  },
};
</script>


<template>
 <div class="p-4 flex items-center justify-center bg-gray-100 n">
        <form class="w-full max-w-screen-md mx-auto">
            <fieldset class="space-y-6">
                <div class="flex items-center justify-between py-4 border-b border-gray-300 ">
                    <legend class="text-2xl text-gray-700 mx-auto ">Pay</legend>
                </div>
                <div class=" bg-white rounded-xl ">
                    <img :src="pay" alt="QR Code"  v-if="pay" class="w-full rounded-xl"/>
                </div>  
            </fieldset>
        </form>
     </div>
</template>