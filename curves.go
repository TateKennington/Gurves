package main;

import "image";
import "image/png";
import "image/color";
import "os";
import "math";

const xMax float64 = 8;
const yMax float64 = 4;
const xMin float64 = -1;
const yMin float64 = -4;

func worldToImage(x, y float64) (int, int){
	return int(1000*(x-xMin)/(xMax-xMin)), int(1000*(1-(y-yMin)/(yMax-yMin)));
}

func drawPoint(graph *image.RGBA, x, y int, c color.RGBA){
	if x>=0 && x<1000 && y>=0 && y<1000{
		graph.Set(x, y, c);
		//c.G = 200;
		//c.B = 200;
		if x-1>=0 && graph.At(x-1,y) == (color.RGBA{255,255,255,255}){
			graph.Set(x-1, y, c);
		}
		if x+1<1000 && graph.At(x+1,y) == (color.RGBA{255,255,255,255}){
			graph.Set(x+1, y, c);
		}
		if y-1>=0 && graph.At(x,y-1) == (color.RGBA{255,255,255,255}){
			graph.Set(x, y-1, c);
		}
		if y+1<1000 && graph.At(x,y+1) == (color.RGBA{255,255,255,255}){
			graph.Set(x, y+1, c);
		}
		if x-1>=0 && y-1>=0 && graph.At(x-1,y-1) == (color.RGBA{255,255,255,255}){
			graph.Set(x-1, y-1, c);
		}
		if x+1<1000 && y+1<1000 && graph.At(x+1,y+1) == (color.RGBA{255,255,255,255}){
			graph.Set(x+1, y+1, c);
		}
		if x-1>=0 && y+1<1000 && graph.At(x-1,y+1) == (color.RGBA{255,255,255,255}){
			graph.Set(x-1, y+1, c);
		}
		if x+1<1000 && y-1>=0 && graph.At(x+1,y-1) == (color.RGBA{255,255,255,255}){
			graph.Set(x+1, y-1, c);
		}
	}
}

func drawPCurve(graph *image.RGBA, curve func(float64)(float64,float64), alpha float64, beta float64, delta float64){
	for t:=alpha+delta; t<beta; t+=delta{
		p1, p2 := worldToImage(curve(t));
		drawPoint(graph, p1, p2, color.RGBA{255,0,0,255});
	}
}

func drawLCurve(graph *image.RGBA, curve func(float64, float64)(float64), step float64){
	for x:=xMin; x<xMax; x+=step{
		for y:=yMin; y<yMax; y+=step{
			if curve(x,y) >= -0.000001 && curve(x,y)<=0.00001{
				p1, p2 := worldToImage(x,y);
				graph.Set(p1, p2, color.RGBA{0, 0, 255, 255});
			}
		}
	}
}

func main(){
	var graph = image.NewRGBA(image.Rect(0,0,1000, 1000));
	file, _ := os.Create("graph.png");
	for x:=0; x<1000; x++{
		for y:=0; y<1000; y++{
			graph.Set(x, y, color.RGBA{255,255,255,255});
		}
	}
	for x:=0; x<1000; x++{
		_, p2 := worldToImage(0, 0);
		graph.Set(x, p2, color.RGBA{0,0,0,255});
	}
	for y:=0; y<1000; y++{
		p1, _ := worldToImage(0, 0);
		graph.Set(p1, y, color.RGBA{0,0,0,255});
	}
	drawPCurve(graph,func(t float64)(float64, float64){
		return  math.Cos(t)*math.Cos(t)*math.Cos(t)*math.Cos(3*t), math.Cos(t)*math.Cos(t)*math.Cos(t)*math.Sin(3*t);
	}, 0, math.Pi,0.001);
	//drawLCurve(graph, func(x, y float64)(float64){
	//	return (x-1)*(x-1)*(x*x+y*y) - x*x;
	//}, 0.0001);
	png.Encode(file, graph);
	file.Close();
}
