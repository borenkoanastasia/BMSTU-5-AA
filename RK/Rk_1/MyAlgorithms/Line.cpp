#include "Line.hpp" 


Line::Line(): z_intersection(Point(N3D))
{}
Line::~Line()
{}
void Line::addPoint(Point p)
{
    points.append(p);
}

QVector<Point> Line::getPoints(Point p)
{
    return points;
}
void Line::setPointZIntersection(Point p)
{
    z_intersection = p;
}
Point Line::getPointZIntersection()
{
    return z_intersection;
}
void Line::transformPoints(Matrix transform)
{
    Matrix z_m(z_intersection);
    Matrix res = z_m*transform;
    //res.output();
    z_intersection = res.getPoint();
    for (int p_i = 0; p_i < points.size(); p_i++)
    {
        Matrix p_m(points[p_i]);
        res = p_m*transform;
        //res.output();
        points[p_i] = res.getPoint();
    }
}

void Line::getXYminXYmax(double &x_min, double &x_max, double &y_min, double &y_max)
{
    //printf("points.size = %d\n", points.size());
    //printf("point size = %d\n", points[0].getSize());
    x_min = points[0].getCurEl(MY_X);
    x_max = points[0].getCurEl(MY_X);
    y_min = points[0].getCurEl(MY_Y);
    y_max = points[0].getCurEl(MY_Y);
    for (int i = 0; i < points.size(); i++)
    {
        //printf("i = %d\t", i);
        if (points[i].getCurEl(MY_X) < x_min)
        {
            x_min = points[i].getCurEl(MY_X);
        }
        if (points[i].getCurEl(MY_X) > x_max)
        {
            x_max = points[i].getCurEl(MY_X);
        }
        if (points[i].getCurEl(MY_Y) < y_min)
        {
            y_min = points[i].getCurEl(MY_Y);
        }
        if (points[i].getCurEl(MY_Y) > y_max)
        {
            y_max = points[i].getCurEl(MY_Y);
        }
    }
    //printf("\nend\n");
}


bool Line::add_smal_edge(int x_start, int y_start, int x_end, int y_end, double *y_min, double *y_max)
{
    int x_len = (x_end - x_start), y_len = (y_end - y_start), len = fabs(x_len);
    if (len == 0){
        len = 1;
    }

    double dx = (double)x_len/(double)len;
    double dy = (double)y_len/(double)len;
    double x = x_start; 
    double y = y_start;
    bool draw_upper = false;
    bool draw_lower = false;
    int cur_x = -1; 

    for (int i = 0; i < len; i++)
    {    
        int horizont_index = (int)x; 
        if (y_min[horizont_index] <= y_max[horizont_index])
        {
            x+=dx;
            y+=dy;
            continue;
        }
        if (y < y_min[horizont_index]) {
            y_min[horizont_index] = y;
        }
        if (y > y_max[horizont_index]) {
            y_max[horizont_index] = y;
        }
        x+=dx;
        y+=dy;
    }
    return true;
}
void Line::addEdge(double *y_min, double *y_max)
{
    //printf("start draw edge\n");
    for (int i = 0; i < points.size() - 1; i++)
    {
        if (!add_smal_edge(points[i]    .getCurEl(MY_X), 
                  points[i]    .getCurEl(MY_Y), 
                  points[i + 1].getCurEl(MY_X),
                  points[i + 1].getCurEl(MY_Y), 
                  y_min, y_max))
            break;
    }
}

void draw_line(int x_start, int y_start, int x_end, int y_end, double *y_min, double *y_max, MyImage *drawer)
{

    int x_len = (x_end - x_start), y_len = (y_end - y_start), len = fabs(x_len);
    if (fabs(x_len) < fabs(y_len))
        len = fabs(y_len);
    if (len == 0){
        //printf("Draw point");
        len = 1;
        return;
    }

    double dx = (double)x_len/(double)len;
    double dy = (double)y_len/(double)len;
    double x = x_start; 
    double y = y_start;
    bool draw_upper = false;
    bool draw_lower = false;
    int cur_x = -1;
    //printf("draw line: start %d,end %d \n", x_start, x_end);    

    for (int i = 0; i < len; i++)
    {    
        //printf("x = %d, y = %d, cur_x = %d", (int)x, (int)y, cur_x);
        int horizont_index = (int)x; 
        if (cur_x != horizont_index){
            draw_upper = false;
            draw_lower = false;
        }
        cur_x = horizont_index;
        if ((y < y_min[horizont_index]) || draw_lower)
        {
            //printf(" Drawed");
            if (y < y_min[horizont_index]) {
                //printf(" (lower moved)");
                y_min[horizont_index] = y;
            }
            drawer->setPoint(MyPoint(x, y));
            draw_lower=true;
        }
        if ((y > y_max[horizont_index]) || draw_upper)
        {
            //printf(" Drawed");

            if (y > y_max[horizont_index]) {
                y_max[horizont_index] = y;
                //printf(" (upper moved)");
            }

            drawer->setPoint(MyPoint(x, y));

            draw_upper=true;
        }
        //printf("\n");
        x+=dx;
        y+=dy;
    }
}

void Line::draw(double *y_min, double *y_max, bool invert, MyImage *drawer)
{
    //printf("start draw line\n");
    int start = 0, end = points.size() - 1, step = 1;
    /*if (invert)
    {
        start = end;
        end = 0;
        step = -1;
    }*/
    int index_points = start; 
    for (int i = 0; i < points.size() - 1; i++)
    {
        draw_line(points[index_points]    .getCurEl(MY_X), 
                  points[index_points]    .getCurEl(MY_Y), 
                  points[index_points + step].getCurEl(MY_X),
                  points[index_points + step].getCurEl(MY_Y), 
                  y_min, y_max, drawer);
        index_points += step;
    }
}

void Line::output()
{
    printf("line: z_intersection:");
    z_intersection.output();
    printf("points\n");
    for (int i = 0; i < points.size(); i++)
    {
        points[i].output();
    }
    printf("\nend line\n");
}

Point Line::getPoint(int i)
{
    return points[i];
}
int Line::getSize()
{
    return points.size();
}