<script>
import axios from 'axios';

export default {
  data() {
    return {
      order: [],
    };
  },
  created() {
    const orderId = this.$route.params.orderId; 
    axios.get(`http://localhost:8080/order/${orderId}`)
      .then(response => {
        this.order =[response.data];
        console.log(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  },
  methods: {
    submit(){
      this.$router.push({ name: 'pay', params: { orderId: this.orderId } });
    }
  }
};
</script>


<template>
     <div class="p-4 flex items-center justify-center bg-gray-100 n">
        <form class="w-full max-w-screen-md mx-auto">
            <fieldset class="space-y-6">
                <div class="flex items-center justify-between py-4 border-b border-gray-300">
                    <div class="mx-10"></div>
                    <legend class="text-2xl text-gray-700 mx-auto ">Order</legend>
                    <div>
                        <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 border border-red-700 rounded">Cancel</button>
                    </div>
                </div>
                <div v-for="orderItem in order" :key="orderItem.orderId" class="container bg-white rounded">
                  <div class="grid grid-cols-5 items-center justify-center border-b border-gray-300">
                      <div></div> <!-- Empty grid cell for the image column -->
                      <div class="text-xl leading-tight uppercase text-center font-bold">Meal</div>
                      <div class="text-xl leading-tight uppercase text-center font-bold">Price</div>
                      <div class="text-xl leading-tight uppercase text-center font-bold">QTY</div>
                      <div class="text-xl leading-tight uppercase text-center font-bold">subtotal</div>
                  </div>  
                  
                  <div v-for="menuItem in orderItem.menuItem" :key="menuItem.menuId" class="grid grid-cols-5 items-center justify-center border-b border-gray-300">  
                        <img :src="menuItem.Menu.menuimg" class="w-60 h-50 rounded"/>
                        <div class="text-xl leading-tight uppercase text-center">{{ menuItem.Menu.menuname }}</div>
                        <div class="text-xl leading-tight uppercase text-center">${{ menuItem.Menu.price }}</div>
                        <div class="text-xl leading-tight uppercase text-center">{{ menuItem.quantity }}</div>
                        <div class="text-xl leading-tight uppercase text-center">${{ menuItem.subtotal }}</div>
                    </div>
                    <div class="text-2xl leading-tight  text-right mr-7 text-red-500">Total ${{orderItem.totalprice }}</div> 
                </div>   
               
            </fieldset>
            <div class="flex items-center justify-center py-4 ">
                <button @click.prevent="submit" class="w-3/4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 border border-blue-700 rounded ">Check</button>
            </div>    
        </form>
     </div>
</template>