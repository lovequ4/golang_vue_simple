<script>
    import axios from 'axios';
 
    export default {
        data() {
            return {
                menu: [],
                selectedItems: [],
                quantity: 1
            };
        },
        created() {
            axios.get('http://localhost:8080/menu')
            .then(response => {
                this.menu = response.data;
                console.log(response.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        methods: {
            decrementQuantity() {
                if (this.quantity > 1) {
                    this.quantity--;
                }
            },
            incrementQuantity() {
                if (this.quantity < 10){
                   this.quantity ++;
                }
            },
            submitForm() {

                
  
                // //儲存 勾選到的資料
                // const selectedMenuIds = this.selectedItems.map(item => item.menuId);

                // //存到 orderdata
                // const orderData = {
                //     menu: this.menu.filter(item => selectedMenuIds.includes(item.menuId)),
                //     quantity:this.quantity
                // };
                    // Create an array of menuitem objects matching the backend JSON structure
                const menuItems = this.selectedItems.map(item => ({
                    menu: item,
                    quantity: this.quantity
                }));

                const orderData = {
                    menuitem: menuItems
                };

    


                console.log(orderData);

                axios.post('http://localhost:8080/addorder', orderData)
                    .then(response => {
                        console.log('API Response:', response.data);
                        this.orderId = response.data.orderId; 
                        this.$router.push({ name: 'order', params: { orderId: this.orderId } });
                    })
                    .catch(error => {
                        console.error('API Error:', error);
                    });
            
            },
            
        },
        
    };
    
</script>


<template>
   
    <!-- https://codepen.io/azrikahar/pen/zYKyQxw  原始語法-->
    <!-- <body class="p-4 flex items-center justify-center bg-gray-100 min-h-screen">
        <form class="w-full max-w-screen-md mx-auto">
            <fieldset class="space-y-6">
            <div class="flex items-center justify-between py-4 border-b border-gray-300">
                <legend class="text-2xl text-gray-700 mr-4">Choose a meal</legend>
                <a href="#" class="font-medium text-gray-500 hover:text-gray-700">Cancel your meal</a>
            </div>
            <div class="grid sm:grid-cols-4 gap-6">
                <label for="plan-hobby" class="relative flex flex-col bg-white p-5 rounded-lg shadow-md cursor-pointer">
                <span class="font-semibold text-gray-500 leading-tight uppercase mb-3">Hobby</span>
                <span class="font-bold text-gray-900">
                    <span class="text-4xl">1</span>
                    <span class="text-2xl uppercase">GB</span>
                </span>
                <span>
                    <span class="text-xl font-bold text-gray-500">$</span>
                    <span class="text-xl font-bold text-gray-900 -ml-1">5</span>
                    <span class="text-xl font-semibold text-gray-500">/</span>
                    <span class="text-lg font-semibold text-gray-500">mo</span>
                </span>
                <input type="radio" name="plan" id="plan-hobby" value="hobby" class="absolute h-0 w-0 appearance-none" />
                <span aria-hidden="true" class="hidden absolute inset-0 border-2 border-green-500 bg-green-200 bg-opacity-10 rounded-lg">
                    <span class="absolute top-4 right-4 h-6 w-6 inline-flex items-center justify-center rounded-full bg-green-200">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-5 w-5 text-green-600">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                    </span>
                </span>
                </label>
                
            </div>
            </fieldset>
        </form>
    </body> -->

    <!-- <div class="p-4 flex items-center justify-center bg-gray-100 n">
        <form class="w-full max-w-screen-md mx-auto">
            <fieldset class="space-y-6">
            <div class="flex items-center justify-center py-4 border-b border-gray-300">
                <legend class="text-2xl text-gray-700 mr-4">Choose a meal</legend>
            </div>
            <div class="grid-cols-4 sm:grid-cols-4 gap-6">
                <label v-for="menuItem in menu" :key="menuItem.menuId" class="relative flex flex-row bg-white p-5 rounded-lg shadow-md cursor-pointer mb-4">
                
                <input type="checkbox"  :value="menuItem" v-model="selectedItems" class="absolute h-0 w-0 appearance-none" />
                <span aria-hidden="true" class="hidden absolute inset-0 border-2 border-green-500 bg-green-200 bg-opacity-10 rounded-lg">
                    <span class="absolute top-4 right-4 h-6 w-6 inline-flex items-center justify-center rounded-full bg-green-200">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-5 w-5 text-green-600">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                    </span>
                </span>
                <img :src="menuItem.menuimg" class="image-size  " />
                    <div class="flex flex-col justify-center  mx-10 ">
                            <div class="text-xl font-bold text-gray-900 leading-tight uppercase mb-3 mt-1">{{ menuItem.menuname }}</div>
                            <div class=" font-semibold text-gray-700">{{ menuItem.description }}</div>
                            <input type="number" min="1" max="10" class="text-gray-900 w-12 h-8 text-center text-xl  border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"/>
                            <div class="text-xl font-bold text-gray-900 text-end mt-auto ">${{ menuItem.price }}</div>
                            
                    </div>
                </label>
      
            </div>
            </fieldset>
            <div class="flex items-center justify-center py-4 ">
                <button @click.prevent="submitForm" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 border border-blue-700 rounded" type="submit">Next</button>
            </div>
        </form>
    </div> -->
    <div class="p-4 flex items-center justify-center bg-gray-100 n">
        <form class="w-full max-w-screen-md mx-auto">
            <fieldset class="space-y-6">
                <div class="flex items-center justify-center py-4 border-b border-gray-300">
                    <legend class="text-2xl text-gray-700 mr-4">Choose a meal</legend>
                </div>
                <div class="grid-cols-4 sm:grid-cols-4 gap-6">
                    <label v-for="menuItem in menu" :key="menuItem.menuId" class="relative flex flex-row bg-white p-5 rounded-lg shadow-md cursor-pointer mb-4">
                        <input type="checkbox" :value="menuItem" v-model="selectedItems" class="absolute h-0 w-0 appearance-none" />
                        <span aria-hidden="true" class="hidden absolute inset-0 border-2 border-green-500 bg-green-200 bg-opacity-10 rounded-lg">
                            <span class="absolute top-4 right-4 h-6 w-6 inline-flex items-center justify-center rounded-full bg-green-200">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-5 w-5 text-green-600">
                                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                                </svg>
                            </span>
                        </span>
                        <div class="flex flex-col justify-center ">
                            <div class="text-xl font-bold text-gray-900 leading-tight uppercase mb-3 mt-1">{{ menuItem.menuname }}</div>
                            <img :src="menuItem.menuimg" class="w-full" />
                            <div class="font-semibold text-gray-700 mt-5">{{ menuItem.description }}</div>
                            <div class="mt-auto">
                                <div class="flex mt-2">
                                    <div class="font-bold text-xl text-gray-900">Quantity：</div>

                                    <button data-action="decrement" class=" bg-gray-300 text-gray-600 hover:text-gray-700 hover:bg-gray-400 h-full w-20 rounded-l cursor-pointer outline-none" @click.prevent="decrementQuantity">
                                        <span class="m-auto text-2xl font-thin">−</span>
                                    </button>
                                    <input type="number" class="focus:outline-none text-center w-full bg-gray-300 font-semibold text-md hover:text-black focus:text-black  md:text-basecursor-default flex items-center text-gray-700  outline-none" name="custom-input-number" v-model="quantity"/>
                                    <button data-action="increment" class="bg-gray-300 text-gray-600 hover:text-gray-700 hover:bg-gray-400 h-full w-20 rounded-r cursor-pointer " @click.prevent="incrementQuantity">
                                        <span class="m-auto text-2xl font-thin" >+</span>
                                    </button>
                                
                                </div>
                                <div class="text-xl font-bold text-gray-900 mt-2">Price: ${{ menuItem.price }}</div>
                            </div>
                        </div>
                    </label>
                </div>
            </fieldset>
            <div class="flex items-center justify-center py-4">
                <button @click.prevent="submitForm" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 border border-blue-700 rounded" type="submit">Next</button>
            </div>
        </form>
    </div>
</template>

<style>
input[type="checkbox"]:checked + span {
  display: block;
}
.image-size {
  width: 500px;
  height: 300px;
}

input[type='number']::-webkit-inner-spin-button,
input[type='number']::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

</style>