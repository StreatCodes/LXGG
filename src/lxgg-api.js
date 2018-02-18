export const getImages = () => {
	return fetch("/api/images", {
		credentials: "include",
	})
	.then((res) => {
		if(res.status == 200){
			return res.json();
		} else {
			throw new Error(res.body);
		}
	})
}



		// const reqs = [];

		// for(const image of images) {
		// 	let request = fetch(image)
		// 	.then(res => {
		// 		return res.json();
		// 	}).then(data => {
		// 		return data.metadata;
		// 	});

		// 	reqs.push(request);
		// }

		// Promise.all(reqs).then(values => {
		// 	this.images = values;
		// });

		// return fetch('/1.0/images')
		// .then(res => {
		// 	return res.json();
		// }).then(data => {
		// 	return data.metadata;
		// })
		
		// const images = this.images.slice();
		// return images.sort((a,b) => {
		// 	if(a.properties.os > b.properties.os){
		// 		return true
		// 	} else {
		// 		return !(b.properties.os > a.properties.os)
		// 	}
		// });