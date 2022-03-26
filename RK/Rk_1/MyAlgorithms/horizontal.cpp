#include "horizontal.hpp"
/*
double getYf1(double x, double z)
{
    return sin(x + z);
}
double getYf2(double x, double z)
{
    return x*x + z*z;
}
double getYf3(double x, double z)
{
    return x+z;
}*/

parameters make_param(double x_start, double x_end, double x_step, double z_start, double z_end, double z_step)
{
    parameters param;
    param.x.start = x_start;
    param.x.end = x_end;
    param.x.step = x_step;
    param.z.start = z_start;
    param.z.end = z_end;
    param.z.step = z_step;
    return param;
}

QVector<Line> make_lines(parameters param, double (*func)(double, double))
{
    QVector<Line> lines;
    Line l;
    Point p(N3D);
    int d;
    //printf("param.z.start %lf; param.z.end %lf; param.z.step %lf\n", param.z.start, param.z.end, param.z.step);
    //printf("param.x.start %lf; param.x.end %lf; param.x.step %lf\n", param.x.start, param.x.end, param.x.step);
    //scanf("%d", &d);
    for (double z = param.z.start; z <= param.z.end; z+=param.z.step)
    {
        l = Line();
        p.setCurEl(MY_X, 0);
        p.setCurEl(MY_Y, 0);
        p.setCurEl(MY_Z, z);
        l.setPointZIntersection(p);
        for (double x = param.x.start; x <= param.x.end; x+=param.x.step)
        {
            //printf("x = %lf, y = %lf, z = %lf\n", x, func(x, z), z);
            p.setCurEl(MY_X, x);
            p.setCurEl(MY_Y, func(x, z));
            p.setCurEl(MY_Z, z);
            l.addPoint(p);
            //p.output();
            //l.output(); 
            //scanf("%d", &d);
        }
        //l.output();
        lines.append(l);
        //lines[lines.size()-1].output();
        //scanf("%d", &d);
    }
    return lines;
}

void transform_lines(QVector<Line> &lines, Matrix transform_matrix)
{
    transform_matrix.output();
    for (int i = 0; i < lines.size(); i++)
    {
       
        Line &l = lines[i];
        //l.output();
        l.transformPoints(transform_matrix);
        //l.output();
        //int d;
        //scanf("%d", &d);
    }
}
void scale_lines(QVector<Line> &lines, Point scale_point, Point center)
{
    Matrix m(M3D, M3D);
    m.make_scale_matrix(scale_point);
    //m.add_center(center);
    //m.output();

    transform_lines(lines, m);
    //lines[0].output();
}
void rotate_lines(QVector<Line> &lines, Point rotate_point)
{
    Matrix m(M3D, M3D);
    m.make_rotate_matrix(rotate_point);

    transform_lines(lines, m);
}
void transfer_lines(QVector<Line> &lines, Point transfer_point)
{
    Matrix m(M3D, M3D);
    m.make_transfer_matrix(transfer_point);
    //m.output();

    transform_lines(lines, m);
}

Point get_scale_point(QVector<Line> lines, double &x, double &y, double &dx, double &dy, double width, double height)
{
    double x_min, x_max, y_min, y_max, cur_x_min, cur_x_max, cur_y_min, cur_y_max;
    lines[0].getXYminXYmax(x_min, x_max, y_min, y_max);
    for (int i = 0; i < lines.size(); i++)
    {
        //printf("i = %d\n", i);
        lines[i].getXYminXYmax(cur_x_min, cur_x_max, cur_y_min, cur_y_max);
        if (cur_x_min < x_min)
            x_min = cur_x_min;
        if (cur_x_max > x_max)
            x_max = cur_x_max;
        if (cur_y_min < y_min)
            y_min = cur_y_min;
        if (cur_y_max > y_max)
            y_max = cur_y_max;
    }
    //printf("x_min = %lf, x_max = %lf, y_min = %lf, y_max = %lf\n", x_min, x_max, y_min, y_max);
    Point ans(N3D);
    x = x_min;
    y = y_max;
    double scale = (width)/(x_max - x_min);
    dx = x_max - x_min;
    dy = y_min - y_max;
    if (scale >(height)/(y_max - y_min))
    {
        scale = (height)/(y_max - y_min);
    }
    ans.setCurEl(MY_X,  scale*0.9);
    ans.setCurEl(MY_Y, -scale*0.9);
    ans.setCurEl(MY_Z, 1);
    return ans;
};

Point get_rotate_point(double x, double y, double z)
{
    Point ans(N3D);
    ans.setCurEl(MY_X, x);
    ans.setCurEl(MY_Y, y);
    ans.setCurEl(MY_Z, z);
    return ans;
};

Point get_center_point(MyImage *drawer)
{
    Point p(N3D);
    p.setCurEl(MY_X, drawer->getWidth()/2);
    p.setCurEl(MY_Y, drawer->getHeight()/2);
    p.setCurEl(MY_Z, 0);
    return p;
}

void make_horizonts(double *y_min, double *y_max, int width,int height)
{
    for (int i = 0; i < width; i++)
    {
        y_min[i] = height;
        y_max[i] = 0;
    }
}

Line draw_edge(QVector<Line> lines, bool front)
{
    Line l;
    int index_point;
    for (int i = 0; i < lines.size(); i++)
    {
        if (front)
        {
            index_point = 0;
        }
        else
        {
            index_point = lines[i].getSize() - 1;
        }
        l.addPoint(lines[i].getPoint(index_point));
    }
    return l;
}

void horizont_alg(QVector<Line> lines, MyImage *drawer)
{
    int width = drawer->getWidth();
    double y_min[width];
    double y_max[width];
    make_horizonts(y_min, y_max, drawer->getWidth(), drawer->getHeight());

    int start = 0, end = lines.size() - 1, step = 1, cur_draw_line;
    bool invert = false;
    if (lines[0].getPointZIntersection().getCurEl(MY_Z) > lines[end].getPointZIntersection().getCurEl(MY_Z))
    {
        start = end;
        end = 0;
        step = -1;
        invert = true;
    }
    cur_draw_line = start;
    for (int counter_draw_line = 0; counter_draw_line < lines.size(); counter_draw_line++)
    {
        //printf("start %d,end %d, step %d, i_draw_line %d\n", start, end, step, counter_draw_line);
        int d;
        //scanf("%d", &d);
        lines[cur_draw_line].draw(y_min, y_max, invert, drawer);
        if (counter_draw_line == 0)
        {
            Line l;
            /*if (y_min[0] > y_max[0])
            { 
                l = draw_edge(lines, invert);
                l.addEdge(y_min, y_max);
            }
            if (y_min[width-1] > y_max[width-1])
            { 
                l = draw_edge(lines, invert);
                l.addEdge(y_min, y_max);
            }*/
            if (invert)
            {
                l = draw_edge(lines, !invert);
                l.addEdge(y_min, y_max);
                l = draw_edge(lines, invert);
                l.addEdge(y_min, y_max);
            }
            else
            {
                l = draw_edge(lines, invert);
                l.addEdge(y_min, y_max);
                l = draw_edge(lines, !invert);
                l.addEdge(y_min, y_max);
            }
            
        }
        cur_draw_line+=step;
    }
}

void show_alg(Matrix current_rotate, MyImage *drawer, parameters param, double (*func)(double, double))
{
    //printf("start show\n\n");
    QVector<Line> lines = make_lines(param, func);
    //printf("make lines\n\n");
    Point center_point = get_center_point(drawer);
    //printf("get center\n\n");
    current_rotate.add_center(center_point);
    //printf("add center\n\n");
    transform_lines(lines, current_rotate);
    //printf("transform lines\n\n");
    double x_min, y_min, dx, dy;
    Point scale_point = get_scale_point(lines, x_min, y_min, dx, dy, drawer->getWidth(), drawer->getHeight());
    //scale_point.output();
    //printf("get scale point\n\n");
    scale_lines(lines, scale_point, center_point);
    //printf("make lines\n\n");
    Point transfer_point(N3D);
    printf("dx = %lf dy = %lf\n",dx, dy);
    transfer_point.setCurEl(MY_X, -x_min*scale_point.getCurEl(MY_X) + dx*0.05*scale_point.getCurEl(MY_X));
    transfer_point.setCurEl(MY_Y, -y_min*scale_point.getCurEl(MY_Y) + dy*0.05*scale_point.getCurEl(MY_Y));
    transfer_point.setCurEl(MY_Z, 0);
    transfer_lines(lines, transfer_point);
    for (int i = 0; i < lines.size(); i++)
    {
    //    lines[i].output();
    }
    horizont_alg(lines, drawer);
    //printf("horizont draw\n\n");
}
